package main

import (
	"context"
	"log"

	"github.com/nXnUs25/gRPCMath/core"
	pb "github.com/nXnUs25/gRPCMath/math/proto"
)

type SumCaller interface {
	Sum(int32, int32) int32
}

func (s *MathServer) Sum(cxt context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum func called %v\n", in)

	var sum SumCaller = &core.Maths{}
	result := sum.Sum(in.GetA(), in.GetB())

	return &pb.SumResponse{Result: result}, nil
}
