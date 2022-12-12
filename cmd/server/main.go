package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/wawan93/grpc-memcached/internal/server"
	"github.com/wawan93/grpc-memcached/pkg/proto"
	"github.com/wawan93/grpc-memcached/pkg/storage/inmemory"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	inmemoryStorage := inmemory.New()
	proto.RegisterStorageServer(s, server.NewServer(inmemoryStorage))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
