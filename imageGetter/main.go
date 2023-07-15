package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("API key is \t", apiKey)

	url := "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&camera=fhaz&api_key=" + apiKey
	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(string(body))

}
