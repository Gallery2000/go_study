package main

import (
	"GoStudy/new_helloworld/client_proxy"
	"fmt"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", ":1234")
	var reply string
	err := client.Hello("booby", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
