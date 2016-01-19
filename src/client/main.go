package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"puller"
)

var listen = ":10001"

func main() {
	var err error
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	var transport thrift.TTransport
	transport, err = thrift.NewTSocket(listen)
	if err != nil {
		fmt.Println("error, thrift init!")
		return
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Printf("error %v\n", err)
		return
	}
	client := puller.NewPullerClientFactory(transport, protocolFactory)
	var request puller.Request
	request.UserId = 12398
	request.Payload = "dlrow olleh"
	var response *puller.Response
	response, err = client.Pull(&request)
	if err != nil {
		fmt.Printf("error, response[%v]\n", err)
		return
	}
	fmt.Printf("response:[%v]\n", response)
}
