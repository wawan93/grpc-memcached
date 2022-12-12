package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/wawan93/grpc-memcached/internal/server"
	"github.com/wawan93/grpc-memcached/internal/storage/memcached"
	"github.com/wawan93/grpc-memcached/pkg/proto"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	memcachedStorage, err := memcached.New("localhost:11211")
	if err != nil {
		log.Fatalf("failed to create memcached storage: %v", err)
	}
	srv := server.NewServer(memcachedStorage)
	proto.RegisterStorageServer(s, srv)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
