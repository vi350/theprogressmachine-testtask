package app

import (
	"encoding/json"
	"github.com/vi350/theprogressmachine-testtask/model"
	"io/ioutil"
	"log"
	"net/http"
)

// FetchCatFact is fetching first page with facts about cats
func FetchCatFact() ([]model.CatFact, error) {
	URL := "https://catfact.ninja/facts"
	// make HTTP request and put response in a string
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("request error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// variable with response (catfacts)
	var response model.CatFactsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("unmarshall error: %s", err)
	}
	return response.CatFacts, nil
}
