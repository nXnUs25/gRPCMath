package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCMath/math/proto"
)

func getPrims(c pb.MathServiceClient) {
	log.Println("Primes: ")
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
	log.Println("Sum: ")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 6,
		B: 5,
	})
	if err != nil {
		log.Fatalf("Failed to greeting: %v", err)
	}

	log.Printf("%v", res.Result)
}

func getAVG(c pb.MathServiceClient) {
	log.Println("AVG: ")

	reqs := []*pb.AvgRequest{
		{Number: 3},
		{Number: 2},
		{Number: 4},
		{Number: 1},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error getting avg: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing stream: %v", err)
	}
	log.Printf("AVG: %v\n", res.GetNumber())
}
