package facades

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jailtonjunior94/go-challenge/address/dtos"
	"github.com/jailtonjunior94/go-challenge/address/enviroments"
)

type IViaCepFacade interface {
	Cep(cep string) (response *dtos.ViaCep, err error)
}

type ViaCepFacade struct {
}

func NewViaCepFacade() IViaCepFacade {
	return &ViaCepFacade{}
}

func (f *ViaCepFacade) Cep(cep string) (response *dtos.ViaCep, err error) {
	res, err := http.Get(fmt.Sprintf(enviroments.ViaCepBaseURL, cep))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, err
}
