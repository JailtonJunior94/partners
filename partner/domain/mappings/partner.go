package mappings

import (
	"github.com/jailtonjunior94/go-challenge/partner/domain/entities"
	"github.com/jailtonjunior94/go-challenge/partner/graph/model"
)

func ToPartnerResponse(e *entities.Partner) *model.Partner {
	return &model.Partner{
		ID:          e.ID.Hex(),
		TradingName: e.TradingName,
		OwnerName:   e.OwnerName,
		Document:    e.Document,
		Address: &model.Address{
			Cep:          e.Address.Cep,
			Street:       e.Address.Street,
			Neighborhood: e.Address.Neighborhood,
			City:         e.Address.City,
			Uf:           e.Address.Uf,
			Location: &model.Location{
				Type: e.Address.Location.Type,
				Lat:  e.Address.Location.Lat,
				Lng:  e.Address.Location.Lng,
			},
		},
	}
}

func ToManyPartnerResponse(partners []*entities.Partner) []*model.Partner {
	var response []*model.Partner

	for _, p := range partners {
		res := &model.Partner{
			ID:          p.ID.Hex(),
			TradingName: p.TradingName,
			OwnerName:   p.OwnerName,
			Document:    p.Document,
			Address: &model.Address{
				Cep:          p.Address.Cep,
				Street:       p.Address.Street,
				Neighborhood: p.Address.Neighborhood,
				City:         p.Address.City,
				Uf:           p.Address.Uf,
				Location: &model.Location{
					Type: p.Address.Location.Type,
					Lat:  p.Address.Location.Lat,
					Lng:  p.Address.Location.Lng,
				},
			},
		}
		response = append(response, res)
	}

	return response
}
