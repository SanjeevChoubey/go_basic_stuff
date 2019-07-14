package main

import (
	context "context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	pb "github.com/Sanjeevchoubey/Shippy/consignment-service/proto/consignment"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Dummy Repository latter part this will be removed when we will started using dB
type Repository struct {
	mu           sync.Mutex
	consignments []*pb.Consignment
}

type service struct {
	repo repository
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func (r *Repository) GetAll() []*pb.Consignment {
	return r.consignments
}

// this method will be handled by gRPC
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil

}

// Create a new consignment
func (r *Repository) Create(req *pb.Consignment) (*pb.Consignment, error) {
	r.mu.Lock()
	r.consignments = append(r.consignments, req)
	r.mu.Unlock()
	return req, nil
}

const (
	port = ":50051"
)

func main() {
	repo := &Repository{}

	//Setup gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	// Register our service to grpc server
	// this will tie our implementation with autogenerated interface code
	// for our protobuf defininition
	pb.RegisterShippingServiceServer(s, &service{repo})

	//Register Reflection services on GRPC server
	reflection.Register(s)

	log.Println("Running on Port: ", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)

	}

}
