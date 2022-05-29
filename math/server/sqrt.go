package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nXnUs25/gRPCMath/core"
	pb "github.com/nXnUs25/gRPCMath/math/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SqrtCaller interface {
	Sqrt(int32) float64
}

func (s *MathServer) Sqrt(cxt context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt func called %v\n", in)

	n := in.GetNumber()

	if n < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Recived negative number %v", n))
	}

	var sq SqrtCaller = &core.Maths{}
	result := sq.Sqrt(n)

	return &pb.SqrtResponse{Number: result}, nil
}
