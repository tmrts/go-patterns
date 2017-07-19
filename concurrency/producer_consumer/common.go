package producer_consumer

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"
)

const (
	TimeoutEnqueueTask = 5 * time.Second
	TimeoutRegister    = 5 * time.Second
	TimeoutDial        = 5 * time.Second
	TimeoutShutdown    = 5 * time.Second

	ProcessDuration = 3 * time.Second
)

var (
	ErrorTEnqueueTask = errors.New("Timeout for enqueuing task")
	ErrorTRegister    = errors.New("Timeout for registering")
	ErrorTDial        = errors.New("Timeout for dailing")

	LogInfo  = log.New(os.Stdout, "[Info]  ", log.Ltime)
	LogError = log.New(os.Stderr, "[Error] ", log.Ltime)
)

func Shutdown(address string, isProducer bool) {
	serviceMethod := "Producer.Shutdown"
	if !isProducer {
		serviceMethod = "Consumer.Shutdown"
	}

	err := rpcCall(address, serviceMethod, struct{}{}, &struct{}{})
	if err != nil {
		LogError.Println(err)
	}
}

func rpcCall(address string, serviceMethod string, args interface{}, reply *struct{}) error {
	conn, err := net.DialTimeout("unix", address, TimeoutDial)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := rpc.NewClient(conn)
	err = client.Call(serviceMethod, args, reply)
	if err != nil {
		return err
	}
	return nil
}

func serverAccept(server *rpc.Server, l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			LogError.Println(err)
			return
		}
		go server.ServeConn(conn)
	}
}
