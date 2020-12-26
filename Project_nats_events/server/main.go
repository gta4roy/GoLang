package main

import (
	pb "Project_nats_events/order"
	dbstore "Project_nats_events/store"
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/nats-io/go-nats"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	port      = ":50051"
	aggregate = "Order"
	event     = "OrderCreate"
)

type server struct {
	storage dbstore.OrderStore
}

func (s *server) CreateOrder(ctx context.Context, in *pb.Order) (*pb.OrderResponse, error) {

	s.storage.CreateOrder(in)
	log.Println("Create Order call..  :")
	go publishOrderCreated(in)
	go normalPublishEvent()
	return &pb.OrderResponse{IsSuccess: true}, nil
}

func (s *server) GetOrders(filter *pb.OrderFilter, stream pb.OrderService_GetOrdersServer) error {
	orders := s.storage.GetOrders()
	for _, orders := range orders {
		if err := stream.Send(&orders); err != nil {
			return err
		}
	}
	log.Println("Gets Orders  :")
	return nil
}

func publishOrderCreated(order *pb.Order) {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()

	eventData, _ := json.Marshal(order)
	event := pb.EventStore{
		AggregateId:   order.OrderId,
		AggregateType: aggregate,
		EventId:       "event_1234",
		EventType:     event,
		EventData:     string(eventData),
	}

	subject := "Order.OrderCreated"
	data, _ := proto.Marshal(&event)
	natsConnection.Publish(subject, data)
	log.Println("Published message on suject :" + subject)
}

func normalPublishEvent() {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()

	subjectNotQueue := "Order.TestEvent"
	data := "String message from the client"
	natsConnection.Publish(subjectNotQueue, []byte(data))
	log.Println("Published message on suject :" + subjectNotQueue)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen %v ", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Println("Server listening on the port :" + port)
	s.Serve(lis)
}
