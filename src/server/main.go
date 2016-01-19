package main

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"puller"
)

type Puller struct {
	Count int64
}

func (p *Puller) Pull(request *puller.Request) (*puller.Response, error) {
	p.Count++
	fmt.Printf("recv req count:%d, req message[%v]\n", p.Count, request)
	return &puller.Response{0, "Success"}, nil
}

func main() {
	var listen string = ":10001"

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(listen)
	if err != nil {
		fmt.Println("error, thrift init!")
		return
	}
	handler := &Puller{0}
	processor := puller.NewPullerProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Printf("server started\n")
	server.Serve()
}
