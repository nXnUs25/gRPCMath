package main

import (
	"io"
	"log"

	"github.com/nXnUs25/gRPCMath/core"
	pb "github.com/nXnUs25/gRPCMath/math/proto"
)

type AvgCaller interface {
	AVG(...int64) float64
}

func (s *MathServer) Avg(stream pb.MathService_AvgServer) error {
	var (
		ac  AvgCaller = &core.Maths{}
		res []int64
	)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Average: %v / %v", res, len(res))
			return stream.SendAndClose(&pb.AvgResponse{
				Number: ac.AVG(res...),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		res = append(res, req.GetNumber())
	}
}
