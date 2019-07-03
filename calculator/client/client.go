package main

import (
	"context"
	"fmt"
	"io"
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

	//doAddition(c)

	// Prime number decomposition
	//doPrimeDecomposition(c)

	// Average
	//doComputeAverage(c)

	// Fins Maximum Number
	doFindMax(c)

}
func doFindMax(c pb.CalculatorServiceClient) {
	fmt.Println("In Find Maximum function")
	
	stream, err := c.FindMaximum(context.Background())
	if err != nil{
		log.Fatalf("Error while calling FindMaximum %v", err)
	}
	waitC := make(chan struct{})
	// Send to server 
	go func(){
		requests := []int32{1,5,3,6,2,20}
		for _,req := range requests{
			stream.Send(&pb.FindMaxRequest{
				Number : req,
			})
		}
		stream.CloseSend()
	}()
	// Recieve 

	go func (){
		resAr := []int32{}
		for{
			res, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("Error in recieving %v\n",err)
				break
			}
			resAr = append(resAr,res.Result)
		}

		fmt.Println(resAr)
		// when recieving is done the close the channel
		close(waitC)
	
	}()

	// Blocking till all response recieved
	<-waitC
}


func doComputeAverage(c pb.CalculatorServiceClient) {
	fmt.Println("In Compute Average- Client")
	requests := []*pb.AvgRequest{
		&pb.AvgRequest{
			Number: 10,
		},
		&pb.AvgRequest{
			Number: 20,
		},
		&pb.AvgRequest{
			Number: 30,
		},
		&pb.AvgRequest{
			Number: 40,
		},
		&pb.AvgRequest{
			Number: 50,
		},
	}


	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error in calling ComputeAverage() %v", err)
	}
for _, req := range  requests{
	stream.Send(req)
}

  res,err := stream.CloseAndRecv()
	if err!= nil{
		log.Fatalf("Error in recieving response %v", err)
	}
	fmt.Println(res)
}

func doPrimeDecomposition(c pb.CalculatorServiceClient) {
	fmt.Println("In Prime Number Decomoposition Client")
	req := &pb.PrimeRequest{
		Number: 126676,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while prime decomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error in reaciving stream %v", err)
		}
		fmt.Println(res)
	}
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
