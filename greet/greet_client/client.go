package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SanjeevChoubey/greet/greet_pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I am at client")
	// Client Connecton
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}
	defer cc.Close()
	//Connect
	c := greet_pb.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c greet_pb.GreetServiceClient) {
	fmt.Println("I am in Unary")
	req := &greet_pb.GreetRequest{
		Greeting: &greet_pb.Greeting{
			FirstName: "sanjeev",
			LastName:  " Choubey",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in Greet Function fetching %v", err)
	}
	fmt.Println("Output from Greet funcation is ", res.Result)
}
