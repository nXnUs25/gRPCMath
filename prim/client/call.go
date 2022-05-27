package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCPrims/prim/proto"
)

func getPrims(c pb.PrimServiceClient) {
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
			log.Fatalf("Failed to read stream on %s - $v\n", res, err)
		}
		log.Println(res.Number)
	}
}
