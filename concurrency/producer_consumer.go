package main

import (
	pc "./producer_consumer"
	"strconv"
	"time"
)

func main() {
	// unix socket address
	pAddr := "/tmp/producer"
	cAddr1 := "/tmp/consumer1"
	cAddr2 := "/tmp/consumer2"

	// producer buffer size
	bufSize := 2

	// start producer
	producer := pc.NewProducer(pAddr, bufSize)
	pc.StartProducer(producer)

	// generate tasks
	for i := 0; i < 10; i++ {
		go func(num int) {
			pc.EnqueueTask(pAddr, pc.Task{strconv.Itoa(num)})
		}(i)
	}

	// start three consumers
	consumer1 := pc.NewConsumer(cAddr1)
	pc.StartConsumer(consumer1, pAddr)
	consumer2 := pc.NewConsumer(cAddr2)
	pc.StartConsumer(consumer2, pAddr)

	time.Sleep(30 * time.Second)
	pc.Shutdown(pAddr, true)
	pc.Shutdown(cAddr1, false)
	pc.Shutdown(cAddr2, false)
}
