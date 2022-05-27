package main

import pb "github.com/nXnUs25/gRPCMath/math/proto"

type MathServer struct {
	pb.MathServiceServer
}

var addr string = "0.0.0.0:50051"
