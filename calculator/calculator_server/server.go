package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go_grpc_server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"time"
)

type server struct{}

// Unary
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
_ = grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
_ = grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))
fmt.Printf("Sum function was invoked with %v\n", req)
nums := req.GetNums()
if nums != nil {

res := &calculatorpb.SumResponse{
Result: nums.FirstNum + nums.SecondNum,
}

		return res, nil
	}
	return nil, nil
}

// Server Streaming
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.SumService_PrimeNumberDecompositionServer) error {
	start := time.Now()
	defer log.Printf("It took %v", time.Since(start))

	var factor int32 = 2
	num := req.GetNum()
	for num > 1 {
		if num%factor == 0 {
			fmt.Println("Divisor =", factor)
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: factor,
			}
			err := stream.Send(res)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			num /= factor
		} else {
			factor++
			// fmt.Printf("Divisor was increased to %v\n", factor)
		}
	}
	return nil
}

// Client Streaming
func (*server) ComputeAverage(stream calculatorpb.SumService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked\n")

	index := float32(0)
	result := float32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 	we have finished reading client stream
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: result / index,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		index += 1
		result += float32(req.Num)
	}

	return nil
}

func (*server) FindMaximum(stream calculatorpb.SumService_FindMaximumServer) error {
	fmt.Println("Started FindMaximum")
	max := int32(0)
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while making request: %v", err)
		}
		number := request.GetNum()
		if number > max {
			max = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{Result: max})
			if sendErr != nil {
				log.Fatalf("Error while sending stream %v", err)
				return err
			}
		}
	}

	return nil
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

var (
	http1Port       = flag.Int("http1_port", 9090, "Port to listen with HTTP1.1 with TLS on.")
	http2Port       = flag.Int("http2_port", 50052, "Port to listen with HTTP2 with TLS on.")
	tlsCertFilePath = flag.String("tls_cert_file", "./calculator/cert/server.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "./calculator/cert/server.key", "Path to the private key file.")
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	calculatorpb.RegisterSumServiceServer(grpcServer, &server{})

	websocketOriginFunc := grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	})
	httpOriginFunc := grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	})

	wrappedServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithWebsockets(true),
		httpOriginFunc,
		websocketOriginFunc,
	)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	http1Server := http.Server{
		Addr:    fmt.Sprintf(":%d", *http1Port),
		Handler: http.HandlerFunc(handler),
	}
	http1Server.TLSNextProto = map[string]func(*http.Server, *tls.Conn, http.Handler){} // Disable HTTP2

	http2Server := http.Server{
		Addr:    fmt.Sprintf(":%d", *http2Port),
		Handler: http.HandlerFunc(handler),
	}

	reflection.Register(grpcServer)


	go func() {
		fmt.Println("Run http:1.1 server")
		if err := http1Server.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http1.1 server: %v", err)
		}
	}()

	go func() {
		fmt.Println("Run http:2.0 server")
		if err := http2Server.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	}()

	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,x-grpc-web")
}
