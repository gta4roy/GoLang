package main

import (
	"context"
	"fmt"
	"hello/hellopb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}

func (*server) AddNumbers(ctx context.Context, request *hellopb.CalculatorRequest) (*hellopb.CalculatorResponse, error) {
	ClientId := request.ClientId
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber

	sum := firstNumber + secondNumber

	response := &hellopb.CalculatorResponse{
		ClientId: ClientId,
		Result:   sum,
	}
	return response, nil
}

func (*server) SubtractNumbers(ctx context.Context, request *hellopb.CalculatorRequest) (*hellopb.CalculatorResponse, error) {
	ClientId := request.ClientId
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber

	sum := firstNumber - secondNumber

	response := &hellopb.CalculatorResponse{
		ClientId: ClientId,
		Result:   sum,
	}
	return response, nil
}

func (*server) MultiplyNumbers(ctx context.Context, request *hellopb.CalculatorRequest) (*hellopb.CalculatorResponse, error) {
	ClientId := request.ClientId
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber

	sum := firstNumber * secondNumber

	response := &hellopb.CalculatorResponse{
		ClientId: ClientId,
		Result:   sum,
	}
	return response, nil
}

func (*server) DevideNumbers(ctx context.Context, request *hellopb.CalculatorRequest) (*hellopb.CalculatorResponse, error) {
	ClientId := request.ClientId
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber

	sum := firstNumber / secondNumber

	response := &hellopb.CalculatorResponse{
		ClientId: ClientId,
		Result:   sum,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error %v", err)
	}
	fmt.Println("Server is listening on %v.....", address)
	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
