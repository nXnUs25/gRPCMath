package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCMath/math/proto"
)

func getPrims(c pb.MathServiceClient) {
	stream, err := c.Primaries(context.Background(), &pb.PrimRequest{Number: 120})
	if err != nil {
		log.Fatalf("Failed to calculate prims: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read stream on %s - %v\n", res, err)
		}
		log.Println(res.Number)
	}
}

func getSum(c pb.MathServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 6,
		B: 5,
	})
	if err != nil {
		log.Fatalf("Failed to greeting: %v", err)
	}

	log.Printf("%v", res.Result)
}
