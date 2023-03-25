package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	prefix         = "https://api.mapbox.com/geocoding/v5/mapbox.places/"
	access         = ".json?access_token=%s"
	wordsFormat    = "%%20%s"
	noWordsMessage = "no words provided to search for"
	apiKey         = "pk.eyJ1Ijoiem9uZXYiLCJhIjoiY2xmbzRhYjZ6MHJ1MzNzbnI1ZHBkazh5ZSJ9.q6dxvrNYWsN_aPQtclEP3Q"
)

type Coordinator struct {
	apiKey string
}

type Response struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Prs []Property `json:"properties"`
}

type Property struct {
	Coords Coordinates `json:"coordinates"`
}

type Coordinates struct {
	First  float64 `json:"longitude"`
	Second float64 `json:"latitude"`
}

func (c *Coordinator) Initialize() {
	c.apiKey = os.Getenv(apiKey)
}

type input struct {
	In string `json:"name"`
}

type objs struct {
	AllIn []input `json:"objects"`
}

func (c *Coordinator) Coordinates(s string) (Coordinates, error) {
	var queryParams string
	pattern := regexp.MustCompile(`[ ,]\s`)
	params := pattern.Split(s, -1)
	l := len(params)
	if l == 0 {
		return Coordinates{}, fmt.Errorf(noWordsMessage)
	} else {
		queryParams = params[0]
		for i, _ := range params {
			if i == 0 {
				continue
			}
			queryParams = queryParams + fmt.Sprintf(wordsFormat, params[i])
		}
	}

	url := prefix + queryParams + fmt.Sprintf(access, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var re Response
	fmt.Println(string(body))
	err = json.Unmarshal(body, &re)
	fmt.Println(re)
	return re.Features[0].Prs, nil
}
