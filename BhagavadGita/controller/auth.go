package controller

import (
	entities "BhagavadGita/entitites"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	CLIENT_ID      = "BZZzLc5JDGztEM1xt4r4bpG4EhoKsG28DXvpO9QD"
	CLIENT_SECRET  = "OJ1EBk6lEg8FBueuaVok17co8ktKO01nAUu7Loi4FthYfoWirC"
	GRANT_TYPE     = "client_credentials"
	SCOPE          = "chapter"
	authUrl        = "https://bhagavadgita.io/auth/oauth/token"
)

func getAuthToken()  (entities.ServiceTokenDetailObject, error) {

	var tokenObjet entities.ServiceTokenDetailObject
	log.Println("get token request initiated")

	// generating token with the help of client and client secret
	data := url.Values{}
	data.Add("client_id", CLIENT_ID)
	data.Add("client_secret", CLIENT_SECRET)
	data.Add("grant_type",GRANT_TYPE)
	data.Add("scope", SCOPE)

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(data.Encode()))

	if err != nil {
		log.Printf("error while perparing http request to fetch the token: %v \n", err.Error())
		return tokenObjet, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}

	log.Println("processing http request to get oauth token")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("error while fetching the token: %v \n", err.Error())
		return tokenObjet, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR in reading output of request: %v \n", err.Error())
		return tokenObjet, err
	}

	err = json.Unmarshal(body, &tokenObjet)
	if err != nil {
		log.Printf("ERROR in decoding output: %v \n", err.Error())
		return tokenObjet, err
	}

	log.Printf("token generated: %s \n",tokenObjet.AccessToken)
	//return the token object
	return  tokenObjet, nil
}