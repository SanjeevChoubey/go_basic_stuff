package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/SanjeevChoubey/calculator/pb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	v1 := req.GetCalculator().Operand1
	v2 := req.GetCalculator().Operand2

	sum := v1 + v2

	res := &pb.CalculatorResponse{
		Result: sum,
	}

	return res, nil
}

func (s *server) PrimeNumberDecomposition(req *pb.PrimeRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	n := req.GetNumber()
	var k int32 = 2
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n = n / k
		} else {
			k++
		}
	}
	return nil
}

func (s *server) ComputeAverage(stream pb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("In Compute average")
	var total, count int32

	for {
		avgReq, err := stream.Recv()
		if err == io.EOF {
			avg := total / count
			stream.SendAndClose(&pb.AvgResponse{
				Result: avg,
			})
		}
		if err != nil {
			log.Fatalf("Error in server %v", err)
		}

		total += avgReq.GetNumber()
		count++

	}

	return nil
}

func main() {
	fmt.Println("In server")

	// Create listner
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while listening servere : %v", err)
	}

	// create new server
	s := grpc.NewServer()

	// Register Calculator Service Server
	pb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serveing %v", err)
	}
}
