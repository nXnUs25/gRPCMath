package main

import (
	"log"

	"github.com/nXnUs25/gRPCPrims/core"
	pb "github.com/nXnUs25/gRPCPrims/prim/proto"
)

type PrimeCaller interface {
	PrimariesAI(int) []int
}

func (p *PrimServer) Primaries(in *pb.PrimRequest, stream pb.PrimService_PrimariesServer) error {
	log.Printf("Primaries excuted: %v", in)
	num := int(in.GetNumber())

	var pc PrimeCaller = &core.Primaries{}
	results := pc.PrimariesAI(num)

	for _, v := range results {
		stream.Send(
			&pb.PrimResponse{
				Number: int32(v),
			},
		)
	}

	return nil
}
