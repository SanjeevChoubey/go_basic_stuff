package main


import(
	context "context"
)

type Handler struct{
	repository
	vesselClient vesselProto.VesselServiceClient
}

// Create Consignmnet- 

func (h *Handler) CreateConsignment(ctx context.Context , req *pb.Consignment, res *pb.Response ) error{

	// in order to get the vessel  exists for given specification, we will call 
	// vessel services
	vr, err:= h.vesselClient.FindAvailable(ctx,&vesselProto.Specification{
		MaxWeight : req.Weight,
		Capacity : int32(len(req.Containers)),
	}) 


}