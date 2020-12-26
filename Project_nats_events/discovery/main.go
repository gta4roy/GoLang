package main

import (
	pb "Project_nats_events/order"
	"log"

	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

var orderServiceUri string

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Config file is not found ")
	}

	orderServiceUri = viper.GetString("discovery.orderservice")
}

func main() {
	//Create connection to the nats server
	natsConnection, err := nats.Connect(nats.DefaultURL)

	if err == nil {
		log.Println("Connected to " + nats.DefaultURL)
		natsConnection.Subscribe("Discovery.OrderService", func(m *nats.Msg) {
			orderServiceDiscovery := pb.ServiceDiscovery{OrderServiceUri: orderServiceUri}
			data, err := proto.Marshal(&orderServiceDiscovery)
			log.Println("URL Discovered ", data)
			if err == nil {
				log.Println("Publishing the data on the network ", data)
				natsConnection.Publish(m.Reply, data)
			}
		})

	} else {
		log.Println("Not Connected to " + nats.DefaultURL)
	}

	for {
	}
}
