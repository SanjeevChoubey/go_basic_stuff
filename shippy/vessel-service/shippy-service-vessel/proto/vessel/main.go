package main

import (
	context "golang.org/x/net/context"
)


type Repository interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

type VesselRepository{
	vessels []*pb.Vessel
}


func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error){
	for _,vessel := range repo.vessels{
		if spec.
	}
}

type Service struct {
	repo Repository
}

func (s *Service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	out.vessel, err := s.repo.FindAvailable(in)
	if err != nil{
		return err
	}

	return nil
}

func main() {
	vessels := *pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel")
	)

	srv.Init()

	pb.

}
