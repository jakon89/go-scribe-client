package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"scribe"
	"crypto/tls"
)

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) (*scribe.ScribeClient, error) {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}
	if transport == nil {
		return nil, fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}
	transport = transportFactory.GetTransport(transport)
	if transport == nil {
		return nil, fmt.Errorf("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}

	err = transport.Open()
	if err != nil {
		return nil, err
	}
	//defer transport.Close()

	factory := scribe.NewScribeClientFactory(transport, protocolFactory)
	return factory, nil
}
