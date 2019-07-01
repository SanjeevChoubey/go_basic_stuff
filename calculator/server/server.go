package main

import (
	"context"
	"fmt"
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
