package main

import (

	"log"
	"os"
	context "golang.org/x/net/context"

	pb "github.com/Sanjeevchoubey/Shippy/vessel-service/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

// type Repository interface {
// 	FindAvailable(*pb.Specification) (*pb.Vessel, error)
// }

// type VesselRepository struct {
// 	vessels []*pb.Vessel
// }

// func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
// 	for _, vessel := range repo.vessels {
// 		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
// 			return vessel, nil
// 		}
// 	}
// 	return nil, errors.New("No vessel found for given spec")
// }

type Service struct {
	repo repository
}

// func (s *Service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
// 	vessel, err := s.repo.FindAvailable(in)
// 	if err != nil {
// 		return err
// 	}
// 	out.Vessel = vessel
// 	return nil
// }
const (
	defaultHost = "datastore:27017"
)

func main() {
	// vessels := []*pb.Vessel{
	// 	&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	// }
	// repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == " "{
		uri = defaultHost
	}

client, err := CreateClient(uri)
if err != nil{
	log.Panic(err)
}
defer client.Disconnect(context.TODO())

vesselCollection := client.Database("shippy").Collection("vessel")

repository := &VesselRepository{
	vesselCollection,
}

	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}

}
