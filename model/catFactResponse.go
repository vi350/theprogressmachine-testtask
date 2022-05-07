package model

import "fmt"

// CatFact is a struct for one fact from response
type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// CatFactsResponse is a struct for json unmarshalling
type CatFactsResponse struct {
	CatFacts []CatFact `json:"data"`
}

// Text for printing out one fact
func (c CatFact) Text() string {
	p := fmt.Sprintf("Fact: %s\nLength: %d\n", c.Fact, c.Length)
	return p
}
