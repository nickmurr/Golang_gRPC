package findMaximum

import (
	"bufio"
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"io"
	"log"
	"os"
	"strconv"
)

func FindMaximum(c calculatorpb.SumServiceClient) {

	fmt.Println("Starting to do BiDirectional Streaming")

	// 	Create Stream by invoking a client

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go send(stream, waitc)

	// receive go routine
	go receive(stream, waitc)

	<-waitc
}

func send(stream calculatorpb.SumService_FindMaximumClient, waitc chan struct{}) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "quit" {
			err := stream.CloseSend()
			close(waitc)
			if err != nil {
				panic(err)
			}
			break
		}
		text := msg
		textNum, err := strconv.ParseInt(text, 10, 32)

		err = stream.Send(&calculatorpb.FindMaximumRequest{Num: int32(textNum)})
		if err != nil {
			panic(err)
		}
	}

	// arr := []int32{4, 5, 2, 76, 8, 3, 5, 7, 34, 5, 346, 34, 6, 234, 346, 457, 23425}
	// for _, number := range arr {
	// 	_ = stream.Send(&calculatorpb.FindMaximumRequest{Num: number})
	// }
	// _ = stream.CloseSend()
}

func receive(stream calculatorpb.SumService_FindMaximumClient, waitc chan struct{}) {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Problem while reading server stream %v\n", err)
		}
		maximum := res.GetResult()
		fmt.Printf("Received New max of...%v\n", maximum)

	}
	// close(waitc)
}
