package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/partner/domain/entities"
	"github.com/jailtonjunior94/go-challenge/partner/graph/generated"
	"github.com/jailtonjunior94/go-challenge/partner/graph/model"
)

func (r *mutationResolver) CreatePartner(ctx context.Context, input model.NewPartner) (*model.Partner, error) {
	partner, err := r.PartnerRepository.Add(entities.NewPartner(input.Name))
	if err != nil {
		return nil, err
	}

	return &model.Partner{
		ID:   partner.ID.Hex(),
		Name: partner.Name,
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
			ID:   p.ID.Hex(),
			Name: p.Name,
			Cep:  "",
		}
		response = append(response, res)
	}

	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
