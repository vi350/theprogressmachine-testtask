package app

import (
	"encoding/json"
	"github.com/vi350/theprogressmachine-testtask/model"
	"io/ioutil"
	"log"
	"net/http"
)

// FetchCrypto is fetching 5m candles for last hour
func FetchCrypto(fiat string, crypto string) ([]model.BinanceCandle, error) {
	URL := "https://api.binance.com/api/v3/klines?symbol=" + crypto + fiat + "&interval=5m&limit=12"
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
	// variable with response (candles)
	var response model.BinanceCandlesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("unmarshall error: %s", err)
	}
	return response.BinanceCandles, nil
}
