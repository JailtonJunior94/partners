package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jailtonjunior94/go-challenge/address/pb"
	"github.com/jailtonjunior94/go-challenge/address/services"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	cepService := services.NewCepGrpcServer()
	pb.RegisterCepServiceServer(grpcServer, cepService)

	log.Printf("🚀 gRPC server is running on http://localhost:%v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over %v: %v", port, err)
	}
}
