package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func downloadFile(url string, filePath string) error {
	reponse, err := http.Get(url)
	if err != nil {
		return err
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code : %d", reponse.StatusCode)
	}

	fichierLocal, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fichierLocal.Close()
	var s int64
	s, err = io.Copy(fichierLocal, reponse.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Write %d byte in %s\n", s, filePath)
	return nil
}

func generateFileName(rover string, sol int, camera string, id int) string {
	basePath := "/home/guillaume/Documents/Projet/NASA_ROVER_API_Movie/output/"
	Path := basePath + rover + "/" + fmt.Sprint(sol) + "/" + camera + "/"
	err := os.MkdirAll(Path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return Path + fmt.Sprint(id) + ".jpg"
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

	for key, value := range dataObj.Data {
		File := generateFileName(value.Rover.Name, value.Sol, value.Camera.Name, key)
		DownErr := downloadFile(value.ImgSrc, File)
		if DownErr != nil {
			log.Fatal(DownErr)
		}
		fmt.Printf("%s\n", File)
	}
}
