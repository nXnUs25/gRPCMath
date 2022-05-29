package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCMath/math/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

/*
The client will send a stream of number (1,5,3,6,2,20) and the server will respond with a stream of (1,5,6,20)
*/
func getMax(c pb.MathServiceClient) {
	log.Println("Max:")
	maxes := []int32{}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error getting max: %v", err)
	}

	done := make(chan struct{})

	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while closing stream: %v", err)
				break
			}
			maxes = append(maxes, res.GetMax())
			log.Printf("Max: %v", res.GetMax())
		}
		close(done)
	}()
	<-done
	log.Printf("Max: %v", maxes)
}

func getSqrt(c pb.MathServiceClient, req int32) {
	log.Println("Sqrt: ")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: req,
	})
	if err != nil {
		// ok variable should tell us if its grpc error
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %v", e.Message())
			log.Printf("Code error from server: %v", e.Code())
			log.Printf("With details: %v", e.Proto().Details)
			if e.Code() == codes.InvalidArgument {
				log.Printf("Wrong value been sent to server: %v, %v", e.Code(), req)
			}
		} else {
			log.Fatalf("A non gRPC error occurred: %v", err)
		}
	}

	log.Printf("%v", res.GetNumber())
}
