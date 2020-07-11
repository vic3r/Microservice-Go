package service

import (
	"context"

	pb "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/proto/consignment"
	r "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/repository"
)

type Service struct {
	Repository r.Repository
}

func (svc *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := svc.Repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (svc *Service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments, err := svc.Repository.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Consignments: consignments}, nil
}
