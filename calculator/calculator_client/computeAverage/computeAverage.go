package computeAverage

import (
	"bufio"
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"log"
	"os"
	"strconv"
	"strings"
)

type Num *calculatorpb.ComputeAverageRequest

func Average(c calculatorpb.SumServiceClient) {
	fmt.Println("Streaming to do a Client Streaming RPC...")

	newNumber := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s ", "Numbers")
	num, _, _ := newNumber.ReadLine()
	numStrings := string(num)
	numSlice := strings.Split(numStrings, " ")
	var numArray []*calculatorpb.ComputeAverageRequest
	for _, v := range numSlice {
		n, _ := strconv.ParseFloat(v, 32)
		numArray = append(numArray, &calculatorpb.ComputeAverageRequest{Num: float32(n)})
	}
	fmt.Println(numArray)

	// reqests := []*calculatorpb.ComputeAverageRequest{
	// 	&calculatorpb.ComputeAverageRequest{
	// 		Num: 1,
	// 	},
	// 	&calculatorpb.ComputeAverageRequest{
	// 		Num: 2,
	// 	},
	// 	&calculatorpb.ComputeAverageRequest{
	// 		Num: 3,
	// 	},
	// 	&calculatorpb.ComputeAverageRequest{
	// 		Num: 4,
	// 	},
	// }

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage %v", err)
	}

	for _, req := range numArray {
		if req.Num != 0 {
			fmt.Printf("Sending Num: %v\n", req.Num)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while making request %v", err)
			}
		}

		// time.Sleep(time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while response: %v", err)
	}
	fmt.Print(response.GetResult())
}
