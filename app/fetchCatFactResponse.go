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
		log.Printf("request error: %s\n", err)
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		log.Printf("closing request error: %s", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("to body conversion error: %s\n", err)
		return nil, err
	}
	// variable with response (catfacts)
	var response model.CatFactsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("unmarshall error: %s\n", err)
		return nil, err
	}
	return response.CatFacts, nil
}
