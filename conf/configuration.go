package conf

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Endpoint Endpoint
}

type Endpoint struct {
	Ethereum Ethereum
}

type Ethereum struct {
	Mainnet string
	Ropsten string
}

var configuration Configuration

func init() {
	viper.SetConfigFile(os.Getenv("GOPATH") + "/src/github.com/DE-labtory/zulu/conf/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
