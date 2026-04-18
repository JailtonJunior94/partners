package enviroments

import (
	"log"

	"github.com/spf13/viper"
)

var (
	ViaCepBaseURL  = ""
	GeocodeBaseURL = ""
	GeocodeKey     = ""
)

func NewSetupEnvironments() {
	var err error

	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	ViaCepBaseURL = viper.GetString("viaCep.baseURL")
	GeocodeBaseURL = viper.GetString("geocode.baseURL")
	GeocodeKey = viper.GetString("geocode.key")
}
