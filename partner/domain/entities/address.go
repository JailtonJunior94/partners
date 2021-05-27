package entities

type Address struct {
	Cep          string    `bson:"cep"`
	Street       string    `bson:"street"`
	Neighborhood string    `bson:"neighborhood"`
	City         string    `bson:"city"`
	Uf           string    `bson:"uf"`
	Location     *Location `bson:"location"`
}

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

func NewAddress(cep, street, neighborhood, city, uf string) *Address {
	return &Address{
		Cep:          cep,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		Uf:           uf,
	}
}

func (a *Address) AddLocation(lat, lng float64) {
	location := &Location{
		Type:        "Point",
		Coordinates: []float64{lat, lng},
	}

	a.Location = location
}

func NewLocation(lat, lng float64) *Location {
	return &Location{
		Type:        "Point",
		Coordinates: []float64{lat, lng},
	}
}
