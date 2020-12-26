package main

import (
	"context"
	"fmt"
	"hello/hellopb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Jeremy"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Println("Received response => [%v]", resp.Greeting)

	additionRequest := &hellopb.CalculatorRequest{ClientId: "abhijit ",
		FirstNumber:  45.34,
		SecondNumber: 67.34,
	}
	calculatedResp, _ := client.AddNumbers(context.Background(), additionRequest)
	fmt.Println("recived Addition calculation request Client ID ", calculatedResp.ClientId, " Result :", calculatedResp.Result)

	calculatedResp2, _ := client.SubtractNumbers(context.Background(), additionRequest)
	fmt.Println("recived Subtract calculation request Client ID ", calculatedResp2.ClientId, " Result :", calculatedResp2.Result)

	calculatedResp3, _ := client.MultiplyNumbers(context.Background(), additionRequest)
	fmt.Println("recived Multiply calculation request Client ID ", calculatedResp3.ClientId, " Result :", calculatedResp3.Result)

	calculatedResp4, _ := client.DevideNumbers(context.Background(), additionRequest)
	fmt.Println("recived Division calculation request Client ID ", calculatedResp4.ClientId, " Result :", calculatedResp4.Result)
}
