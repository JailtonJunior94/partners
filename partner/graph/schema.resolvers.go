package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/partner/domain/entities"
	"github.com/jailtonjunior94/go-challenge/partner/domain/mappings"
	"github.com/jailtonjunior94/go-challenge/partner/graph/generated"
	"github.com/jailtonjunior94/go-challenge/partner/graph/model"
)

func (r *mutationResolver) CreatePartner(ctx context.Context, input model.NewPartner) (*model.Partner, error) {
	res, err := r.AddressFacade.Address(input.Cep)
	if err != nil {
		return nil, err
	}

	address := entities.NewAddress(res.Cep, res.Street, res.Neighborhood, res.City, res.Uf)
	address.AddLocation(float64(res.Location.Lat), float64(res.Location.Lng))

	newPartner := entities.NewPartner(input.TradingName, input.OwnerName, input.Document)
	newPartner.AddAddress(address)

	partner, err := r.PartnerRepository.Add(newPartner)
	if err != nil {
		return nil, err
	}

	return mappings.ToPartnerResponse(partner), nil
}

func (r *queryResolver) Partners(ctx context.Context) ([]*model.Partner, error) {
	partners, err := r.PartnerRepository.Get()
	if err != nil {
		return nil, err
	}

	if len(partners) == 0 {
		return make([]*model.Partner, 0), nil
	}

	return mappings.ToManyPartnerResponse(partners), nil
}

func (r *queryResolver) Partner(ctx context.Context, id string) (*model.Partner, error) {
	partner, err := r.PartnerRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if partner == nil {
		return nil, nil
	}

	return mappings.ToPartnerResponse(partner), nil
}

func (r *queryResolver) PartnerByDistance(ctx context.Context, cep string, distance int) ([]*model.Partner, error) {
	res, err := r.AddressFacade.Address(cep)
	if err != nil {
		return nil, err
	}

	location := entities.NewLocation(float64(res.Location.Lat), float64(res.Location.Lng))
	partners, err := r.PartnerRepository.GetByLocation(location, distance)
	if err != nil {
		return nil, err
	}

	if len(partners) == 0 {
		return make([]*model.Partner, 0), nil
	}

	return mappings.ToManyPartnerResponse(partners), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
