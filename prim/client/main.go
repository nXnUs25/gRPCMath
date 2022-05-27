package main

import (
	"log"

	pb "github.com/nXnUs25/gRPCPrims/prim/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %s: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewPrimServiceClient(conn)
	getPrims(c)
}
