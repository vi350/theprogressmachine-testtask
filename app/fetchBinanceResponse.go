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
	// variable with response (candles)
	var response model.BinanceCandlesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("unmarshall error: %s\n", err)
		return nil, err
	}
	return response.BinanceCandles, nil
}
