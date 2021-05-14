package services

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/address/pb"
)

type CepGrpcServer struct {
	pb.UnimplementedCepServiceServer
}

func (c *CepGrpcServer) GetCep(ctx context.Context, in *pb.CepRequest) (*pb.CepResponse, error) {
	return &pb.CepResponse{
		Cep: "06503015",
	}, nil
}

func NewCepGrpcServer() *CepGrpcServer {
	return &CepGrpcServer{}
}
