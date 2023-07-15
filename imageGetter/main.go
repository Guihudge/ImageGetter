package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	conf "guihudge.com/ImageGetter/config"
)

func GetApiKeyFromConfig(config string) (string, error) {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.SetConfigType("yml")
	var configuration conf.Configuration

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return "", err
	}

	return configuration.ApiKey, nil
}

func main() {
	var apiKey string
	var err error

	apiKey, err = GetApiKeyFromConfig("config")
	if err != nil {
		fmt.Print("Reead config error:", err)
		os.Exit(1)
	}

	fmt.Println("API key is \t", apiKey)

}
