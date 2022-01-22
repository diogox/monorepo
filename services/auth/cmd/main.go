package main

import (
	"fmt"
	"log"

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
	defer conn.Close()

	var (
		protocolFactory  = frugal.NewFProtocolFactory(thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{}))
		transportFactory = frugal.NewFNatsSubscriberTransportFactory(conn)
		provider         = frugal.NewFScopeProvider(nil, transportFactory, protocolFactory)
		subscriber       = event.NewEventsSubscriber(provider)
	)

	sub, err := subscriber.SubscribeEventCreated(func(ctx frugal.FContext, ev *event.Event) {
		fmt.Println(ev.GetID())
		fmt.Println(ev.GetMessage())
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	wait := make(chan bool)
	log.Println("Subscriber started...")
	<-wait
}
