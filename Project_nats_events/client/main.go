package main

import (
	pb "Project_nats_events/order"
	"context"
	"io"
	"log"
	"time"

	"github.com/nats-io/go-nats"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func createOrders(client pb.OrderServiceClient) {
	order := &pb.Order{
		OrderId:   "Order_1234",
		Status:    "Created",
		CreatedOn: time.Now().Unix(),
		OrderItems: []*pb.Order_OrderItem{
			&pb.Order_OrderItem{
				Code:      "knd100",
				Name:      "Kindle Voyage",
				UnitPrice: 220,
				Quantity:  1,
			},
			&pb.Order_OrderItem{

				Code:      "kc101",
				Name:      "Kindle Voyage SmartShell Case",
				UnitPrice: 10,
				Quantity:  2,
			},
		},
	}
	resp, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		log.Fatalf("Could not create order: %v", err)
	}
	if resp.IsSuccess {
		log.Printf("A new Order has been placed with id: %s", order.OrderId)
	} else {
		log.Printf("Error:%s", resp.Error)
	}
}

func getOrders(client pb.OrderServiceClient, filter *pb.OrderFilter) {
	// calling the streaming API
	stream, err := client.GetOrders(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get Orders: %v", err)
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetOrders(_) = _, %v", client, err)
		}
		log.Printf("Order: %v", order)
	}
}
func main() {
	//create NATS server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()

	log.Println("Connected to " + nats.DefaultURL)
	msg, err := natsConnection.Request("Discovery.OrderService", nil, 1000*time.Millisecond)
	if err == nil && msg != nil {

		log.Println("Sending request to discovery")
		orderServiceDiscovery := pb.ServiceDiscovery{}
		err := proto.Unmarshal(msg.Data, &orderServiceDiscovery)
		if err != nil {
			log.Fatal("Error in unmarhsall %v", err)
		}
		address := orderServiceDiscovery.OrderServiceUri

		log.Println("Order service endpoint found ", address)

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatal("could not connect to grpc service ")
		}
		defer conn.Close()
		client := pb.NewOrderServiceClient(conn)

		createOrders(client)

		filter := &pb.OrderFilter{SearchText: ""}
		log.Println("------------Order--------------------")
		getOrders(client, filter)
	} else {
		log.Println("-Unable to send the request due to disconnection -")
	}

}
