package main

import (
	"fmt"
	"go_grpc_server/calculator/calculator_client/primeNumber"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, i'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewSumServiceClient(cc)

	/* SUM */
	// newChannel := make(chan int32)
	// for {
	// 	go sum.Sum(c, newChannel)
	// 	fmt.Printf("Sum = %v\n\n", <-newChannel)
	// }

	/* Prime NUMBER */
	primeNumber.PrimeNumber(c)
}
