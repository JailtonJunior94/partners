package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jailtonjunior94/go-challenge/address/enviroments"
	"github.com/jailtonjunior94/go-challenge/address/facades"
	"github.com/jailtonjunior94/go-challenge/address/pb"
	"github.com/jailtonjunior94/go-challenge/address/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort = "9000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	enviroments.NewSetupEnvironments()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	viaCepFacade := facades.NewViaCepFacade()
	geoCodeFacade := facades.NewGeocodeFacade()
	cepService := services.NewCepGrpcServer(viaCepFacade, geoCodeFacade)

	pb.RegisterCepServiceServer(grpcServer, cepService)

	log.Printf("🚀 gRPC server is running on http://localhost:%v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over %v: %v", port, err)
	}
}
