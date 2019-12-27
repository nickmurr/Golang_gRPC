package main

import (
	"context"
	"fmt"
	"go_grpc_server/greet/greetpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("Greet function was invoked with %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := fmt.Sprintf("Hello %s %s", firstName, lastName)

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTime func was invoked with %v\n", req)

	firstName := req.Greeting.GetFirstName()

	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s, number %v", firstName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		err := stream.Send(res)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
