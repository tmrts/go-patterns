package producer_consumer

import (
	"net"
	"net/rpc"
	"os"
	"time"
)

type Producer struct {
	address string
	tasks   chan Task
	done    chan struct{}

	consumers chan string
	l         net.Listener
}

func NewProducer(address string, capacity int) *Producer {
	return &Producer{
		address:   address,
		tasks:     make(chan Task, capacity),
		consumers: make(chan string, capacity),
		done:      make(chan struct{}, 0),
	}
}

func StartProducer(p *Producer) {
	server := rpc.NewServer()
	server.Register(p)

	os.Remove(p.address)
	l, err := net.Listen("unix", p.address)
	if err != nil {
		LogError.Println(err)
	}

	p.l = l
	go serverAccept(server, p.l)
	go p.schedule()
}

func EnqueueTask(pAddr string, task Task) {
	err := rpcCall(pAddr, "Producer.Enqueue", task, &struct{}{})
	if err != nil {
		LogError.Println(err)
	}
}

func (p *Producer) schedule() {
	for {
		select {
		case task := <-p.tasks:
			// one consumer consumes one job at one time
			consumer := <-p.consumers

			go func(consumer string) {
				// if rpcCall fails, this consumer will be regarded as unavailable.
				err := rpcCall(consumer, "Consumer.DoTask", task, &struct{}{})
				if err != nil {
					LogError.Printf("[%s] %s", p.address, err.Error())
				} else { // re-register consumer
					go func(consumer string) {
						select {
						case p.consumers <- consumer:
						case <-time.After(5 * time.Second):
							LogError.Printf("[%s] %s", p.address, ErrorTRegister)
						}
					}(consumer)
				}
			}(consumer)
		case <-p.done:
			break
		}
	}
}

func (p *Producer) Enqueue(task Task, _ *struct{}) error {
	select {
	case p.tasks <- task:
		LogInfo.Printf("[%s] enqueue task[%s] successfully!", p.address, task.String())
	case <-time.After(TimeoutEnqueueTask):
		return ErrorTEnqueueTask
	}

	return nil
}

func (p *Producer) Register(address string, _ *struct{}) error {
	select {
	case p.consumers <- address:
		LogInfo.Printf("[%s] register consumer[%s] successfully!", p.address, address)
	case <-time.After(TimeoutRegister):
		return ErrorTRegister
	}

	return nil
}

func (p *Producer) Shutdown(_ struct{}, _ *struct{}) error {
	p.done <- struct{}{}
	if err := p.l.Close(); err != nil {
		return err
	}

	LogInfo.Printf("[%s] shutdown successfully!\n", p.address)
	return nil
}
