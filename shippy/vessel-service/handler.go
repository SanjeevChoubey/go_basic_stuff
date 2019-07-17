package main

import (
  context "context"
  pb "github.com/Sanjeevchoubey/Shippy/vessel-service/shippy-service-vessel/proto/vessel"
)

type handler struct{
    repository
}

func (s *handler) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	vessel, err := s.repository.FindAvailable(in)
	if err != nil {
		return err
	}
	out.Vessel = vessel
	return nil
}

func (s *handler) Create(ctx context.Context, in *pb.Vessel, out *pb.Response) error {
  err := s.repository.Create(in)
  if err != nil{
    return err
  }
  return nil
}
