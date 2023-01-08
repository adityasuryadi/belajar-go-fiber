package config

import (
	"encoding/json"
	"fmt"
	"go-blog/model"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ConfigGoogle() *oauth2.Config {
	// var configuration Config
	// config := &oauth2.Config{
	// 	ClientID:     configuration.Get("OAUTH_CLIENT"),
	// 	ClientSecret: configuration.Get("OAUTH_SECRET"),
	// 	RedirectURL:  configuration.Get("OAUTH_REDIRECT_URL"),
	// 	Scopes: []string{
	// 		"https://www.googleapis.com/auth/userinfo.email"}, // you can use other scopes to get more data
	// 	Endpoint: google.Endpoint,
	// }
	// return config
	config := &oauth2.Config{
		ClientID:     "401541805025-ren6oshuflcp4qrn6uhoagg1tbkb3gts.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:3001/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email"}, // you can use other scopes to get more data
		Endpoint: google.Endpoint,
	}
	return config
}

func GetClient(token string) model.GoogleResponse {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken}},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)

	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var data model.GoogleResponse
	errorz := json.Unmarshal(body, &data)
	if errorz != nil {

		panic(errorz)
	}
	return data
}
