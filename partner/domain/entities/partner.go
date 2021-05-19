package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Partner struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func NewPartner(name string) *Partner {
	return &Partner{
		ID:   primitive.NewObjectID(),
		Name: name,
	}
}
