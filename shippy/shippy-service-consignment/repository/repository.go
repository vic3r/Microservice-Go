package repository

import (
	"context"

	pb "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/proto/consignment"
)

type Repository interface {
	Create(context.Context, *pb.Consignment) (*pb.Consignment, error)
	GetAll(context.Context, *pb.GetRequest) ([]*pb.Consignment, error)
}
