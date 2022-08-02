package main

import (
	v1 "GoStudy/stream_grpc_test/proto/common/stream/proto/v1"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

type server struct {
	v1.UnimplementedGreeterServer
}

func (s *server) GetStream(req *v1.StreamReqData, res v1.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&v1.StreamResData{
			Data: fmt.Sprint("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PutStream(cliStr v1.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *server) AllStream(allStr v1.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			a, _ := allStr.Recv()
			fmt.Println("收到客户端消息" + a.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&v1.StreamResData{Data: "我是服务端"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	v1.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
