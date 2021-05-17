package facades

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jailtonjunior94/go-challenge/address/dtos"
	"github.com/jailtonjunior94/go-challenge/address/enviroments"
)

type IGeocodeFacade interface {
	GeoCodeByAddress(cep string) (response *dtos.Geocode, err error)
}

type GeocodeFacade struct {
}

func NewGeocodeFacade() IGeocodeFacade {
	return &GeocodeFacade{}
}

func (f *GeocodeFacade) GeoCodeByAddress(address string) (response *dtos.Geocode, err error) {
	res, err := http.Get(fmt.Sprintf(enviroments.GeocodeBaseURL, url.QueryEscape(address), enviroments.GeocodeKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, err
}
