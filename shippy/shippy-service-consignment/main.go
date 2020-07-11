package main

import (
	"log"
	"net"

	pb "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/proto/consignment"
	repository "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/repository/storage"
	svc "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	repository := &repository.Storage{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service := &svc.Service{Repository: repository}
	pb.RegisterShippingServiceServer(s, service)

	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
