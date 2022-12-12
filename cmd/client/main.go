package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/wawan93/grpc-memcached/pkg/proto"
)

var (
	addr   = flag.String("addr", "localhost:8080", "the address to connect to")
	method = flag.String("method", "get", "the method to call")
	key    = flag.String("key", "", "the key")
	value  = flag.String("value", "", "the value")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}()
	c := proto.NewStorageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *method {
	case "get":
		r, err := c.Get(ctx, &proto.GetRequest{Key: *key})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("get response: %s", r.String())
	case "set":
		r, err := c.Set(ctx, &proto.SetRequest{Key: *key, Body: *value})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("set response: %s", r.String())
	case "delete":
		r, err := c.Delete(ctx, &proto.DeleteRequest{Key: *key})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("delete response: %s", r.String())
	}
}
