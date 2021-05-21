package repositories

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/partner/domain/entities"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/database"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/environments"

	"go.mongodb.org/mongo-driver/bson"
)

type IPartnerRepository interface {
	Add(p *entities.Partner) (partner *entities.Partner, err error)
	Get() (partners []*entities.Partner, err error)
}

type PartnerRepository struct {
	Client database.IMongoConnection
}

func NewPartnerRepository(client database.IMongoConnection) IPartnerRepository {
	return &PartnerRepository{Client: client}
}

func (r *PartnerRepository) Add(p *entities.Partner) (partner *entities.Partner, err error) {
	collection, err := r.Client.GetCollection(environments.Database, environments.Collection)
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(context.Background(), &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PartnerRepository) Get() (partners []*entities.Partner, err error) {
	collection, err := r.Client.GetCollection(environments.Database, environments.Collection)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &partners); err != nil {
		return nil, err
	}

	return partners, nil
}
