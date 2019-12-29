package main

import (
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
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

func (*server) FindMaximum(stream calculatorpb.SumService_FindMaximumServer) error {
	fmt.Println("Started FindMaximum")
	max := int32(0)
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while making request: %v", err)
		}
		number := request.GetNum()
		if number > max {
			max = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{Result: max})
			if sendErr != nil {
				log.Fatalf("Error while sending stream %v", err)
				return err
			}

		}
	}

	return nil
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
