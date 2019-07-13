package main

import (
	"sync"
	context "context"
)

type repository interface{
	Create()
}
type Repository struct{
	mu sync.Mutex
	consignments []*pb.Consignment
}

type service struct{
	repo repository
}
func (s *service) CreateConsignment(context.Context, *pb.Consignment) (*Response, error) {

}

func main() {

}
