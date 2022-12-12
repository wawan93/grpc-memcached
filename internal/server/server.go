package server

import (
	"context"

	"github.com/wawan93/grpc-memcached/pkg/proto"
)

type Server struct {
	proto.UnimplementedStorageServer

	storage Storage
}

func NewServer(storage Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (res *proto.GetResponse, err error) {
	var value string
	value, err = s.storage.Get(ctx, req.Key)
	if err != nil {
		return
	}

	res = &proto.GetResponse{
		Body: value,
	}
	return
}

func (s *Server) Set(ctx context.Context, req *proto.SetRequest) (*proto.SetResponse, error) {
	err := s.storage.Set(ctx, req.Key, req.Body)
	return &proto.SetResponse{}, err
}

func (s *Server) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	err := s.storage.Delete(ctx, req.Key)
	return &proto.DeleteResponse{}, err
}
