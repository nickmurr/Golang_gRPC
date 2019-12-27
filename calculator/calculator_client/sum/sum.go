package sum

import (
	"bufio"
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"log"
	"os"
	"strconv"
)

type myInt []byte
type myReader bufio.Reader

var num myReader

func Sum(c calculatorpb.SumServiceClient, newChan chan int32) {

	firstNum := num.ToReader("first")
	secondNum := num.ToReader("second")

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

func (m myInt) ToInt() int32{
	f := string(m)
	num, _ := strconv.ParseInt(f, 10, 32)

	return int32(num)
}

func (m myReader) ToReader(label string) int32 {
	var num myInt

	newNumber := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s number: ", label)
	num, _, _ = newNumber.ReadLine()

	return num.ToInt()
}
