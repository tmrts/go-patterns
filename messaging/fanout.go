package messaging

import (
	"context"
	"go.uber.org/zap"
	"sync"
	"sync/atomic"
)

var (
	log, _ = zap.NewDevelopment()
)

const (
	MaxWorkers      = 16
	MaxQueueSize    = 512
	MasterQueueSize = MaxQueueSize * MaxWorkers
)

type Pipeline struct {
	workers map[int]*worker
	chain   chan interface{}
}

func (p *Pipeline) Start() {
	go func(pipe *Pipeline) {
		for {
			expectationWorkers := len(pipe.chain)
			if expectationWorkers > MaxWorkers {
				expectationWorkers = expectationWorkers % MaxWorkers
			}
			for _, c := range pipe.workers {
				if expectationWorkers < int(c.index) {
					break
				}
				select {
				case val := <-pipe.chain:
					go c.stream(val)
				}
			}
		}
	}(p)
}

func (p *Pipeline) Dispatch(msg interface{}) {
	p.chain <- msg
}

type DispatcherBuilder func() Dispatcher

func NewPipeline(d DispatcherBuilder, idle uint32, debug bool) *Pipeline {
	ch := make(chan interface{}, MasterQueueSize)
	wk := make(map[int]*worker)
	for i := 0; i < MaxWorkers; i++ {
		wk[i] = &worker{
			index:      uint32(i + 1),
			chain:      make(chan interface{}, MaxQueueSize),
			mutex:      new(sync.Mutex),
			debug:      debug,
			idle:       idle,
			Dispatcher: d(),
		}
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
		ctx, cancel := context.WithCancel(context.Background())
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
		err := c.Before(ctx)

		if err != nil {
			log.Error("can not start worker", zap.Error(err))
		}
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
				}
			default:
				atomic.AddUint32(&idle, 1)
				if i := atomic.LoadUint32(&idle); i > 0 {
					if i > c.idle {
						return
					}
				}
			}
		}
	}
}
