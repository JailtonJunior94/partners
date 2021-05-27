package graph

import (
	"github.com/jailtonjunior94/go-challenge/partner/facades"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/database"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/repositories"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MongoConnection   database.IMongoConnection
	PartnerRepository repositories.IPartnerRepository
	AddressFacade     facades.IAddressFacade
}

func NewResolver() *Resolver {
	connection := database.NewConnection()
	connection.CreateIndex()
	return &Resolver{
		MongoConnection:   connection,
		PartnerRepository: repositories.NewPartnerRepository(connection),
		AddressFacade:     facades.NewAddressFacade(),
	}
}
