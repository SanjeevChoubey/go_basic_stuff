package main

import (
	context "context"

	pb "github.com/Sanjeevchoubey/Shippy/consignment-service/proto/consignment"
	vesselProto "github.com/Sanjeevchoubey/Shippy/vessel-service/shippy-service-vessel/proto/vessel"
)

type Handler struct {
	repository
	vesselClient vesselProto.VesselServiceClient
}

// Create Consignmnet-

func (h *Handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// in order to get the vessel  exists for given specification, we will call
	// vessel services
	vr, err := h.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil{
		return err
	}

	req.VesselId = vr.Vessel.Id

	// save consignment

	if err := h.repository.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil

}

func (h *Handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := h.repository.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
