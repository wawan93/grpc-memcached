package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/wawan93/grpc-memcached/config"
	"github.com/wawan93/grpc-memcached/internal/server"
	"github.com/wawan93/grpc-memcached/internal/storage/inmemory"
	"github.com/wawan93/grpc-memcached/internal/storage/memcached"
	memcachedClient "github.com/wawan93/grpc-memcached/pkg/memcached"
	"github.com/wawan93/grpc-memcached/pkg/proto"
)

const (
	inMemoryStorageType  = "inmemory"
	memcachedStorageType = "memcached"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	client, err := memcachedClient.NewMemcached(cfg.Memcached.Addr)
	if err != nil {
		log.Fatalf("failed to connect to memcached: %v", err)
	}

	var storage server.Storage

	if cfg.Storage == inMemoryStorageType {
		storage = inmemory.New()
	}

	if cfg.Storage == memcachedStorageType {
		storage = memcached.New(client)
		if err != nil {
			log.Fatalf("failed to create memcached storage: %v", err)
		}
	}

	srv := server.NewServer(storage)
	proto.RegisterStorageServer(s, srv)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
