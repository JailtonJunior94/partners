package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jailtonjunior94/go-challenge/partner/domain/entities"
	"github.com/jailtonjunior94/go-challenge/partner/graph/generated"
	"github.com/jailtonjunior94/go-challenge/partner/graph/model"
)

func (r *mutationResolver) CreatePartner(ctx context.Context, input model.NewPartner) (*model.Partner, error) {
	address := entities.NewAddress("06503-015", "Rua José Pontes Zé Buraco", "Parque Fernão Dias", "Santana de Parnaíba", "SP")
	address.AddCoordinates(-23.455110549926758, -46.92776107788086)

	newPartner := entities.NewPartner(input.TradingName, input.OwnerName, input.Document)
	newPartner.AddAddress(address)

	partner, err := r.PartnerRepository.Add(newPartner)
	if err != nil {
		return nil, err
	}

	return &model.Partner{
		ID:          partner.ID.Hex(),
		TradingName: partner.TradingName,
	}, nil
}

func (r *queryResolver) Partners(ctx context.Context) ([]*model.Partner, error) {
	partners, err := r.PartnerRepository.Get()
	if err != nil {
		return nil, err
	}

	if len(partners) == 0 {
		return make([]*model.Partner, 0), nil
	}

	var response []*model.Partner
	for _, p := range partners {
		res := &model.Partner{
			ID:          p.ID.Hex(),
			TradingName: p.TradingName,
			Address:     &model.Address{},
		}
		response = append(response, res)
	}

	return response, nil
}

func (r *queryResolver) Partner(ctx context.Context, id string) (*model.Partner, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
