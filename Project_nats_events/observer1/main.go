package main

import (
	pb "Project_nats_events/order"
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
	"google.golang.org/protobuf/proto"
)

const (
	queue   = "Order.OrdersCreatedQueue"
	subject = "Order.OrderCreated"

	subjectNotQueue = "Order.TestEvent"
)

func main() {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			log.Println("Subscribe message in worker 1 %v \n", eventStore)
		}
	})

	natsConnection.Subscribe(subjectNotQueue, func(msg *nats.Msg) {
		log.Println("Subscribe message in worker 1 %v \n", msg)
	})

	runtime.Goexit()
}
