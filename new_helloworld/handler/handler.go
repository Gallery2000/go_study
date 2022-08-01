package handler

const HelloServiceName = "handler/HelloService"

type NewHelloService struct {
}

func (s *NewHelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
