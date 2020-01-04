package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	crudserver "go_grpc_server/blog/blog_server/crud"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_grpc_server/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var (
	http1Port       = flag.Int("http1_port", 9090, "Port to listen with HTTP1.1 with TLS on.")
	http2Port       = flag.Int("http2_port", 50052, "Port to listen with HTTP2 with TLS on.")
	tlsCertFilePath = flag.String("tls_cert_file", "./calculator/cert/server.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "./calculator/cert/server.key", "Path to the private key file.")
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(grpcServer, &crudserver.CrudServer{})

	websocketOriginFunc := grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	})

	httpOriginFunc := grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	})

	fmt.Printf("Connecting to MongoDB\n\n")

	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	crudserver.Collection = client.Database("mydb").Collection("blog")

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

	go func() {
		fmt.Println("Start TCP Server is running on port :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Printf("All in run\n\n")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal a received
	<-ch

	fmt.Println("Stopping a server")
	grpcServer.Stop()
	fmt.Println("Closing a listener")
	_ = lis.Close()
	fmt.Println("Closing mongodb Connection")
	_ = client.Disconnect(context.TODO())
	fmt.Println("End of Program")
}
