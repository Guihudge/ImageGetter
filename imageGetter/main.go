package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	conf "guihudge.com/ImageGetter/config"
	datatype "guihudge.com/ImageGetter/dataType"
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

func generateApiUrl(sol int, rover string, camera string, apiKey string) string {
	baseUrl := "https://api.nasa.gov/mars-photos/api/v1/rovers/"
	url := baseUrl + rover + "/photos?sol=" + fmt.Sprint(sol) + "&camera=" + camera + "&api_key=" + apiKey
	fmt.Printf("Url: %s", url)
	return url
}

func extractDataFromUrl(url string) (datatype.Data, error) {
	resp, getErr := http.Get(url)
	dataObj := datatype.Data{}

	if getErr != nil {
		return dataObj, getErr
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return dataObj, readErr
	}

	jsonErr := json.Unmarshal(body, &dataObj)
	if jsonErr != nil {
		return dataObj, jsonErr
	}

	return dataObj, nil
}

func main() {
	var apiKey string
	var err error

	apiKey, err = GetApiKeyFromConfig("config")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	url := generateApiUrl(999, "curiosity", "fhaz", apiKey)

	dataObj, err := extractDataFromUrl(url)

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range dataObj.Data {
		fmt.Println(value.Rover.Name)
		fmt.Println(value.ImgSrc)
		println("")
	}
}
