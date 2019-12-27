package primeNumber

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

type myInt []byte
type myReader bufio.Reader

var num myReader

func PrimeNumber(c calculatorpb.SumServiceClient)  {
	numb := num.ToReader("initial")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Num: numb,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
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

func (m myReader) ToReader(label string) int32 {
	var num myInt

	newNumber := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s number: ", label)
	num, _, _ = newNumber.ReadLine()

	return num.ToInt()
}

func (m myInt) ToInt() int32 {
	f := string(m)
	num, _ := strconv.ParseInt(f, 10, 32)

	return int32(num)
}
