package environments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	ConnectionString = ""
	Database         = ""
	Collection       = ""
	AddressBaseURL   = ""
)

func NewSetupEnvironments() {
	var err error

	viper.SetConfigName(fmt.Sprintf("settings.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	ConnectionString = viper.GetString("mongo.connectionString")
	Database = viper.GetString("mongo.database")
	Collection = viper.GetString("mongo.collection")
	AddressBaseURL = viper.GetString("address.baseURL")
}
