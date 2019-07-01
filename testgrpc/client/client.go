package main

import (
	"fmt"
	"log"

	"github.com/SanjeevChoubey/testgrpc/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %v:", err)
	}
	defer cc.Close()

	c := pb.NewEmployeeServiceClient(cc)
	fmt.Println("Created client", c)
}
