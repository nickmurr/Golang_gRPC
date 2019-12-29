package squreRoot

import (
	"context"
	"fmt"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func SquareRoot(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// incorrect call
	doErrorCall(c, -10)
}

func doErrorCall(c calculatorpb.SumServiceClient, number int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: int32(number)})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// 	Actual error from gRPC (user error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number")
			}
		} else {
			log.Fatalf("Big err calling Square Root %v", err)
		}
		return
	}
	fmt.Printf("Result of square root of %v is %v\n\n", number, res.GetNumberRoot())

}
