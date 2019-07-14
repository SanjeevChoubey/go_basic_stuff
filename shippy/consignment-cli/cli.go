//shippy/consignment-cli

package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "github.com/Sanjeevchoubey/Shippy/consignment-service/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFileName = "consignment.json"
)

func parseFile(f string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Error reading json %v", err)
	}

	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {
	//setup connection to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting service %v", err)
	}
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	// Get the input from json
	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Error parsing file %v", err)
	}
	res, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Error parsing file %v", err)
	}

	res, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Error parsing file %v", err)
	}
	for _, val := range res.Consignments {
		log.Println(val)
	}
}
