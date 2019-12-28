package main

import (
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

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

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	nums := req.GetNums()

	res := &calculatorpb.SumResponse{
		Result: nums.FirstNum + nums.SecondNum,
	}

	return res, nil
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
