package main

import (
	"fmt"
	"log"
	"net"

	"context"

	"github.com/SanjeevChoubey/greet/greet_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greet_pb.GreetRequest) (*greet_pb.GreetResponse, error) {
	fmt.Printf("Greet Function was invoked %v", req)
	firstName := req.GetGreeting().FirstName
	result := " Hello " + firstName
	res := &greet_pb.GreetResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("Hello to Server!!!")

	// Listner
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen :%v", err)
	}

	// Create new  grpc Server
	s := grpc.NewServer()

	greet_pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failes to serve %v: ", err)
	}

}
