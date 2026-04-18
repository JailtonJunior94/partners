package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Partner struct {
	ID          primitive.ObjectID `bson:"_id"`
	TradingName string             `bson:"tradingName"`
	OwnerName   string             `bson:"ownerName"`
	Document    string             `bson:"document"`
	Address     *Address           `bson:"address"`
}

func NewPartner(tradingName, ownerName, document string) *Partner {
	return &Partner{
		ID:          primitive.NewObjectID(),
		TradingName: tradingName,
		OwnerName:   ownerName,
		Document:    document,
	}
}

func (p *Partner) AddAddress(address *Address) {
	p.Address = address
}
