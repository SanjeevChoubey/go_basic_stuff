package main

import (
	context "golang.org/x/net/context"
	
)


type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct{
	vessels []*pb.Vessel
}


func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error){
	for _,vessel := range repo.vessels{
		
	}
}

type Service struct {
	repo Repository
}

func (s *Service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	vessel, err := s.repo.FindAvailable(in)
	if err != nil{
		return err
	}
	out.veseel = vessel
	return nil
}

func main() {
	vessels := *pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(),&service{repo})

	if err := srv.Run(); err != nil{
		fmt.println(err)
	}

}
