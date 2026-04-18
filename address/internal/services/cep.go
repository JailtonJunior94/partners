package services

import (
	"context"

	"github.com/jailtonjunior94/partners/address/internal/facades"
	"github.com/jailtonjunior94/partners/address/internal/validators"
	"github.com/jailtonjunior94/partners/proto/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CepGrpcServer struct {
	pb.UnimplementedCepServiceServer
	ViaCep  facades.IViaCepFacade
	Geocode facades.IGeocodeFacade
}

func NewCepGrpcServer(c facades.IViaCepFacade, g facades.IGeocodeFacade) *CepGrpcServer {
	return &CepGrpcServer{ViaCep: c, Geocode: g}
}

func (c *CepGrpcServer) GetCep(ctx context.Context, in *pb.CepRequest) (*pb.CepResponse, error) {
	cepValidated := validators.RemoveSpecialCharacters(in.Cep)
	if cepValidated == "" {
		return nil, status.Errorf(codes.InvalidArgument, "O CEP é obrigatório")
	}

	address, err := c.ViaCep.Cep(in.Cep)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	geoCode, err := c.Geocode.GeoCodeByAddress(address.Street)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if len(geoCode.Results) == 0 {
		return nil, status.Errorf(codes.NotFound, "%v", err)
	}

	return &pb.CepResponse{
		Cep:          address.Cep,
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		Uf:           address.Uf,
		Location: &pb.LocationResponse{
			Lat: geoCode.Results[0].Geometry.Location.Lat,
			Lng: geoCode.Results[0].Geometry.Location.Lng,
		},
	}, nil
}
