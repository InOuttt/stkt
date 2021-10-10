package movie

import (
	"context"

	pb "github.com/inouttt/stkt/test_2"
)

type Grpc struct {
	pb.UnimplementedMovieServiceServer
	svc Services
}

func NewGrpc(svc Services) *Grpc {
	return &Grpc{svc: svc}
}

func (s *Grpc) Search(ctx context.Context, req *pb.SearchRequest) (*pb.Response, error) {
	sr := SearchRequest{
		Apikey: req.ApiKey,
		Search: req.ApiKey,
		Page:   req.ApiKey,
	}

	rec, err := s.svc.Search(sr)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Payload: []byte(rec),
	}, nil
}
func (s *Grpc) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.Response, error) {
	dr := DetailRequest{
		Apikey: req.ApiKey,
		Id:     req.Id,
	}

	rec, err := s.svc.Detail(dr)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Payload: []byte(rec),
	}, nil
}
