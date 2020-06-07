package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/EwanValentine/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/EwanValentine/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Set-up micro instance
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.client", service.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
