package main

import (
	"log"

	"github.com/nXnUs25/gRPCMath/core"
	pb "github.com/nXnUs25/gRPCMath/math/proto"
)

type PrimeCaller interface {
	PrimariesAI(int) []int
}

func (p *MathServer) Primaries(in *pb.PrimRequest, stream pb.MathService_PrimariesServer) error {
	log.Printf("Primaries excuted: %v", in)
	num := int(in.GetNumber())

	var pc PrimeCaller = &core.Maths{}
	results := pc.PrimariesAI(num)

	for _, v := range results {
		stream.Send(
			&pb.PrimResponse{
				Number: int64(v),
			},
		)
	}

	return nil
}
