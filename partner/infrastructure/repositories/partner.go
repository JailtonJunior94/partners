package repositories

import (
	"context"

	"github.com/jailtonjunior94/go-challenge/partner/domain/entities"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/database"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/environments"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPartnerRepository interface {
	Add(p *entities.Partner) (partner *entities.Partner, err error)
	Get() (partners []*entities.Partner, err error)
	GetById(id string) (partner *entities.Partner, err error)
	GetByLocation(location *entities.Location, distance int) (partners []*entities.Partner, err error)
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

func (r *PartnerRepository) GetById(id string) (partner *entities.Partner, err error) {
	collection, err := r.Client.GetCollection(environments.Database, environments.Collection)
	if err != nil {
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&partner); err != nil {
		return nil, nil
	}

	return partner, nil
}

func (r *PartnerRepository) GetByLocation(location *entities.Location, distance int) (partners []*entities.Partner, err error) {
	collection, err := r.Client.GetCollection(environments.Database, environments.Collection)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"address.location",
			bson.D{
				{"$near", bson.D{
					{"$geometry", location},
					{"$maxDistance", distance},
				}},
			}},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &partners); err != nil {
		return nil, err
	}

	return partners, nil
}
