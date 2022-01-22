package main

import (
	"github.com/Workiva/frugal/lib/go"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/nats-io/nats.go"

	"github.com/diogox/monorepo/pkg/contracts/gen-go/event"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	var (
		protocolFactory  = frugal.NewFProtocolFactory(thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{}))
		transportFactory = frugal.NewFNatsPublisherTransportFactory(conn)
		provider         = frugal.NewFScopeProvider(transportFactory, nil, protocolFactory)
		publisher        = event.NewEventsPublisher(provider)
	)
	if err := publisher.Open(); err != nil {
		panic(err)
	}
	defer publisher.Close()

	err = publisher.PublishEventCreated(frugal.NewFContext("1"), &event.Event{
		ID:      3,
		Message: "Hello World",
	})
	if err != nil {
		panic(err)
	}

	for {
	}
}
