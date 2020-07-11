package storage

import (
	"context"
	"sync"

	pb "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/proto/consignment"
	r "github.com/vic3r/Microservice-Go/shippy/shippy-service-consignment/repository"
)

type Storage struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

var _ r.Repository = &Storage{}

func (sto *Storage) Create(ctx context.Context, consignment *pb.Consignment) (*pb.Consignment, error) {
	sto.mu.Lock()
	updated := append(sto.consignments, consignment)
	sto.consignments = updated
	sto.mu.Unlock()

	return consignment, nil
}

func (sto *Storage) GetAll(ctx context.Context, req *pb.GetRequest) ([]*pb.Consignment, error) {
	return sto.consignments, nil
}
