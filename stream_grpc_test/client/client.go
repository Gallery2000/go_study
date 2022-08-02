package main

import (
	v1 "GoStudy/stream_grpc_test/proto/common/stream/proto/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := v1.NewGreeterClient(conn)

	//res, _ := c.GetStream(context.Background(), &v1.StreamReqData{Data: "慕课网"})
	//for {
	//	a, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(a.Data)
	//}
	//
	//putS, _ := c.PutStream(context.Background())
	//i := 0
	//for {
	//	i++
	//	putS.Send(&v1.StreamReqData{
	//		Data: fmt.Sprintf("啦啦啦%d", i),
	//	})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}

	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			a, _ := allStr.Recv()
			fmt.Println("收到服务端消息" + a.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&v1.StreamReqData{Data: "我是客户端"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
