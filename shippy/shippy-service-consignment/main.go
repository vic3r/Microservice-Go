package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/proto/consignment"
	repository "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/repository/storage"
	svc "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/service"
)

const (
	port = ":50051"
)

func main() {
	repository := &repository.Storage{}

	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	service.Init()
	if err := pb.RegisterShippingServiceHandler(service.Server(), &svc.Service{Repository: repository}); err != nil {
		log.Panic(err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
