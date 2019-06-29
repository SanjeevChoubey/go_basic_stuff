package main

import (
	"context"
	"log"
	"net"

	"github.com/SanjeevChoubey/testgrpc/pb"
	"github.com/SanjeevChoubey/testgrpc/testgrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const port = "9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	pb.RegisterEmployeeServiceServer(s, new(employeeService))
	log.Println("Starting server on port:", port)
	s.Serve(lis)
}

type employeeService struct{}

func (e *employeeService) GetAllEmployees(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	return nil, nil
}
