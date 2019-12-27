package main

import (
	"bufio"
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

type myInt []byte
type myReader bufio.Reader

var num myReader

func main() {
	fmt.Println("Hello, i'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewSumServiceClient(cc)
	newChannel := make(chan int64)

	for {
		go Sum(c, newChannel)
		fmt.Printf("Sum = %v\n\n", <-newChannel)
	}

}

func Sum(c calculatorpb.SumServiceClient, newChan chan int64) {

	firstNum := num.toReader("first")
	secondNum := num.toReader("second")

	req := &calculatorpb.SumRequest{
		Nums: &calculatorpb.Nums{
			FirstNum:  firstNum,
			SecondNum: secondNum,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	newChan <- res.Result

}

func (m myInt) toInt() int64 {
	f := string(m)

	num, _ := strconv.ParseInt(f, 10, 64)

	return num
}

func (m myReader) toReader(label string) int64 {
	var num myInt

	newNumber := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter %s number: ", label)

	num, _, _ = newNumber.ReadLine()

	return num.toInt()
}
