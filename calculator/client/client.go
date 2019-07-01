package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SanjeevChoubey/calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	// Client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error in client connection %v", err)
	}

	defer cc.Close()

	//Connect
	c := pb.NewCalculatorServiceClient(cc)

	doAddition(c)

}

func doAddition(c pb.CalculatorServiceClient) {
	fmt.Println("In doAddition")
	req := &pb.CalculatorRequest{
		Calculator: &pb.Calculator{
			Operand1: 3,
			Operand2: 10,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling sum %v", err)
	}
	fmt.Println("sum of 2 numbers 3 & 10 is ", res.Result)
}
