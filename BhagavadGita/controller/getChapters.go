package controller

import (
	entities "BhagavadGita/entitites"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var getChaptersUrl = "https://bhagavadgita.io/api/v1/chapters?access_token="
var fileLocation = "C:\\BhagavadGita\\test.json"

func GetChapters(_ time.Time) error{

	log.Println("Get chapters request initiated")

	//get oauth2 token
	token, err := getAuthToken()
	if err != nil {
		log.Printf("error while fetching oauth token: %v \n",err.Error())
		return err
	}

	accessToken := token.AccessToken

	client := &http.Client{}
	req, err := http.NewRequest("GET", getChaptersUrl+accessToken, nil)
	if err != nil {
		log.Printf("error while perparing http request to fetch the chapters: %v \n", err.Error())
		return err
	}

	req.Header.Add("Accept", "application/json")
	log.Println("processing http request to get chapters")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error while getting chapters: %v \n", err.Error())
		return  err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR in reading output of request: %v \n", err.Error())
		return err
	}

	var responseObject []entities.ChapterSchema
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		log.Printf("ERROR in decoding output: %v \n", err.Error())
		return err
	}

	//writing json response to a file
	file, _ := json.MarshalIndent(responseObject, "", " ")

	_ = ioutil.WriteFile(fileLocation, file, 0644)

	log.Println("response stored successfully in json file")
	fmt.Printf("Json response stored successfully at location %v \n", fileLocation)

	return nil
}
