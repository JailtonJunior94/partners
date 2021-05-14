package main

import (
	"log"
	"net"

	"github.com/jailtonjunior94/go-challenge/address/pb"
	"github.com/jailtonjunior94/go-challenge/address/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	cepService := services.NewCepGrpcServer()
	pb.RegisterCepServiceServer(grpcServer, cepService)

	log.Println("🚀 gRPC server is running on http://localhost:9000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over 9000: %v", err)
	}
}
