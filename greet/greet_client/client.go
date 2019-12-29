package main

import (
	"context"
	"fmt"
	"go_grpc_server/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello, i'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDiStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Niko",
			LastName:  "Muraviov",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Niko",
			LastName:  "Muraviov",
		},
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break //  we've reached end of the stream
		}
		if err != nil {
			log.Fatalf("Error while reading a stream: %v", err)
		}
		log.Printf("Response from Greet: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Streaming to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Niko",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephan",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mark",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Diana",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}
	// Iterate over slice and send each message individual
	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req.Greeting.FirstName)
		err := stream.Send(req)

		if err != nil {
			log.Fatalf("Error while making request: %v", err)
		}
		time.Sleep(time.Second)

	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while response: %v", err)
	}
	fmt.Println(response.Result)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do BiDirectional Streaming")

	// we create a stream by invoking a client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Niko",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephan",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mark",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Diana",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	}

	waitc := make(chan struct{})

	// we send a bunch of messages to the client (go routine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending Message %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending request: %v", err)
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error while closing stream")
		}
	}()

	// we receive a bunch of messages from the client (go routine)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while sending request: %v", err)
			}
			fmt.Printf("Received: %s\n", msg.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
