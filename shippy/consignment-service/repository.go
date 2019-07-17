package main

import (
	context "context"

	pb "github.com/Sanjeevchoubey/Shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

//Create a new consignment
func (r *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := r.collection.InsertOne(context.Background(), consignment)
	return err
}

func (r *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	row, err := r.collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}
	var consignments []*pb.Consignment
	for row.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := row.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}

	return consignments, nil
}
