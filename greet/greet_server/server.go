package main

import (
	"context"
	"fmt"
	"go_grpc_server/greet/greetpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// 	client cancelled the request
			fmt.Println("client cancelled the request!")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}
		time.Sleep(time.Second)
	}

	firstName := req.GetGreeting().FirstName
	result := "Hello " + firstName
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

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
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet func was invoked with a streaming request\n")
	result := "Hello "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 	we have finished reading client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstname := req.GetGreeting().GetFirstName()
		result += firstname + "! "

	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone func was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading a stream: %v,", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := fmt.Sprintf("Hello %s!", firstName)

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v,", err)
			return sendErr
		}
	}
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"

	creed, sslErr := credentials.NewClientTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failed loading certificates %v", sslErr)
		return
	}

	s := grpc.NewServer(grpc.Creds(creed), grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))

	reflection.Register(s)

	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
