package facades

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/address/pb"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/environments"
	"google.golang.org/grpc"
)

type IAddressFacade interface {
	Address(cep string) (response *pb.CepResponse, err error)
}

type AddressFacade struct {
}

func NewAddressFacade() IAddressFacade {
	return &AddressFacade{}
}

func (f *AddressFacade) Address(cep string) (response *pb.CepResponse, err error) {
	connection, err := grpc.Dial(environments.AddressBaseURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	client := pb.NewCepServiceClient(connection)
	request := pb.CepRequest{Cep: cep}

	res, err := client.GetCep(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
