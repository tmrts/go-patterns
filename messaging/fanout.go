package messaging

import (
	"context"
	"go.uber.org/zap"
	"io"
	"sync"
	"sync/atomic"
)

const (
	MaxWorkers   = 32
	MaxQueueSize = 128
)

var (
	running uint32 = 0
	log, _         = zap.NewDevelopment()
)

type Pipeline struct {
	workers []*worker
	chain   chan interface{}
}

func (p *Pipeline) Start() {
	distributeToChannels := func(ch chan interface{}, cs []*worker) {
		writer := cs[0] //first worker must stream as default
		for {
			for _, c := range cs {
				expectationWorkers := uint32(len(ch)/(MaxQueueSize/MaxWorkers)) + 1
				runningWorker := atomic.LoadUint32(&running)
				if c.index > runningWorker+1 || c.index > expectationWorkers {
					break
				}
				select {
				case val := <-ch:
					writer = c
					if c.debug {
						log.Info("Worker receiving", zap.Any("index", writer.index), zap.Any("running", runningWorker), zap.Any("no# workers", expectationWorkers))
					}
					go writer.stream(val)
				}
			}
			}
		}
	}

	go distributeToChannels(p.chain, p.workers)
}

func (p *Pipeline) Dispatch(msg interface{}) {
	p.chain <- msg
}

type DispatcherBuilder func() Dispatcher

func NewPipeline(d DispatcherBuilder, idle uint32, debug bool) *Pipeline {
	wk := make([]*worker, 0, MaxWorkers)
	ch := make(chan interface{}, 4096)
	for i := 0; i < MaxWorkers; i++ {
		wk = append(wk,
			&worker{
				index:      uint32(i + 1),
				chain:      make(chan interface{}, MaxQueueSize),
				mutex:      new(sync.Mutex),
				debug:      debug,
				idle:       idle,
				Dispatcher: d(),
			})
	}
	return &Pipeline{workers: wk, chain: ch}
}

type Dispatcher interface {
	Before(context.Context) error
	After() error
	Process(interface{}) error
}

type worker struct {
	index   uint32
	mutex   *sync.Mutex
	running bool
	chain   chan interface{}
	debug   bool
	idle    uint32
	Dispatcher
}

func (c *worker) stream(val interface{}) {
	c.chain <- val
	if !c.running {
		c.mutex.Lock()
		c.running = true
		atomic.AddUint32(&running, 1)
		defer atomic.AddUint32(&running, ^uint32(1-1))
		ctx, cancel := context.WithCancel(context.Background())
		err := c.Before(ctx)

		if err != nil {
			log.Error("can not start worker", zap.Error(err))
		}
		defer func(w *worker, cancel context.CancelFunc) {
			if w.debug {
				log.Info("Worker leaving", zap.Any("index", w.index), zap.Any("idle", w.idle))
			}
			err := c.After()
			if err != nil {
				log.Error("can not finish track issue", zap.Error(err))
			}
			cancel()
			w.mutex.Unlock()
			w.running = false
		}(c, cancel)
		var idle uint32 = 0
		for {
			select {
			case msg := <-c.chain:
				atomic.StoreUint32(&idle, 0)
				if msg != nil {
					err := c.Process(msg)
					if err != nil {
						log.Error("can not process message",
							zap.Any("msg", &msg),
							zap.Error(err),
						)
					}
					if err == io.EOF {
						return
					}
				}
			default:
				atomic.AddUint32(&idle, 1)
				if i := atomic.LoadUint32(&idle); i > 0 {
					if i > c.idle {
						return
					}
					if c.debug {
						log.Info("Idle", zap.Any("worker index", c.index), zap.Any("idle", idle))
					}
				}
			}
		}
	}
}
