package main

import (
	"io"
	"log"

	"github.com/nXnUs25/gRPCMath/core"
	pb "github.com/nXnUs25/gRPCMath/math/proto"
)

type MaxCaller interface {
	Max(int32, int32) int32
}

func (s *MathServer) Max(stream pb.MathService_MaxServer) error {
	var (
		mc  MaxCaller = &core.Maths{}
		max int32     = 0
	)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading stream: \n%v", err)
		}
		log.Printf("Max: %v - Num: %v", max, req.GetNumber())
		res := mc.Max(max, req.GetNumber())
		if res == max {
			continue
		}
		max = res
		err = stream.Send(&pb.MaxResponse{
			Max: max,
		})

		if err != nil {
			log.Fatalf("Error while sending response to client: \n%v", err)
		}
	}
}
