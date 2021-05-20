package entities

type Address struct {
	Cep          string    `bson:"cep"`
	Street       string    `bson:"street"`
	Neighborhood string    `bson:"neighborhood"`
	City         string    `bson:"city"`
	Uf           string    `bson:"uf"`
	Type         string    `bson:"type"`
	Coordinates  []float64 `bson:"coordinates"`
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

func (a *Address) AddCoordinates(lat, lng float64) {
	a.Type = "Point"
	a.Coordinates = []float64{lat, lng}
}
