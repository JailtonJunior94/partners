package services

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/address/facades"
	"github.com/jailtonjunior94/go-challenge/address/pb"
)

type CepGrpcServer struct {
	pb.UnimplementedCepServiceServer
	ViaCep facades.IViaCepFacade
}

func NewCepGrpcServer(c facades.IViaCepFacade) *CepGrpcServer {
	return &CepGrpcServer{ViaCep: c}
}

func (c *CepGrpcServer) GetCep(ctx context.Context, in *pb.CepRequest) (*pb.CepResponse, error) {
	address, err := c.ViaCep.Cep(in.Cep)
	if err != nil {
		return nil, err
	}

	return &pb.CepResponse{
		Cep:          address.Cep,
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
	}, nil
}
