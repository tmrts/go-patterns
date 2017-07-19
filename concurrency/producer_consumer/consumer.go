package producer_consumer

import (
	"net"
	"net/rpc"
	"os"
	"time"
)

type Consumer struct {
	address string

	l net.Listener
}

func NewConsumer(address string) *Consumer {
	return &Consumer{address: address}
}

func StartConsumer(c *Consumer, pAddr string) {
	server := rpc.NewServer()
	server.Register(c)

	os.Remove(c.address)
	l, err := net.Listen("unix", c.address)
	if err != nil {
		LogError.Println(err)
	}

	c.l = l
	go serverAccept(server, c.l)

	err = rpcCall(pAddr, "Producer.Register", c.address, &struct{}{})
	if err != nil {
		LogError.Println(err)
		c.l.Close()
	}
}

func (c *Consumer) DoTask(task Task, _ *struct{}) error {
	// takes time to do the task
	time.Sleep(ProcessDuration)

	LogInfo.Printf("[%s] finish task[%s] successfully!", c.address, task.String())
	return nil
}

func (c *Consumer) Shutdown(_ struct{}, _ *struct{}) error {
	if err := c.l.Close(); err != nil {
		return err
	}

	LogInfo.Printf("[%s] shutdown successfully!", c.address)
	return nil
}
