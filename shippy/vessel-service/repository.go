package main

import(
  "context"
  "go.mongodb.org/mongo-driver/mongo"
	
	"gopkg.in/mgo.v2/bson"
  pb "github.com/Sanjeevchoubey/Shippy/vessel-service/shippy-service-vessel/proto/vessel"
)
type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
  Create(in *pb.Vessel) error
}

type VesselRepository struct {
	vessels *mongo.Collection
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
  filter := bson.D{{
  		"capacity",
  		bson.D{{
  			"$lte",
  			spec.Capacity,
  		}, {
  			"$lte",
  			spec.MaxWeight,
  		}},
  	}}
 var vessel *pb.Vessel
 if err := repo.vessels.FindOne(context.TODO(),filter).Decode(&vessel); err != nil{
   return nil, err
 }
	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error{
   _,err := repo.vessels.InsertOne(context.TODO(), vessel)
   if err != nil{
     return err
   }
   return nil
}
