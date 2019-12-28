package main

import (
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct{}

// Unary
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	nums := req.GetNums()

	res := &calculatorpb.SumResponse{
		Result: nums.FirstNum + nums.SecondNum,
	}

	return res, nil
}

// Server Streaming
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.SumService_PrimeNumberDecompositionServer) error {
	start := time.Now()
	defer log.Printf("It took %v", time.Since(start))

	var factor int32 = 2
	num := req.GetNum()
	for num > 1 {
		if num%factor == 0 {
			fmt.Println("Divisor =", factor)
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: factor,
			}
			err := stream.Send(res)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			num /= factor
		} else {
			factor++
			// fmt.Printf("Divisor was increased to %v\n", factor)
		}
	}
	return nil
}

// Client Streaming
func (*server) ComputeAverage(stream calculatorpb.SumService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked\n")

	index := float32(0)
	result := float32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 	we have finished reading client stream
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: result / index,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		index += 1
		result += float32(req.Num)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterSumServiceServer(s, &server{})

	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
