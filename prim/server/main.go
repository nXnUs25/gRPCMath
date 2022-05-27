package main

import (
	"log"
	"net"

	pb "github.com/nXnUs25/gRPCPrims/prim/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}
	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterPrimServiceServer(s, &PrimServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s: %v", addr, err)
	}
}
