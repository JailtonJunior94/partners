package structs

type ViaCep struct {
	Cep          string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	Uf           string `json:"uf"`
}
