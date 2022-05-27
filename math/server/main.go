package main

import (
	"log"
	"net"

	pb "github.com/nXnUs25/gRPCMath/math/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}
	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterMathServiceServer(s, &MathServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s: %v", addr, err)
	}
}
