package app

import (
	"fmt"
	"log"
)

func BaseLevel() error {
	candles, err := FetchCrypto("USDT", "BTC")
	if err != nil {
		log.Printf("fetch error: %s\n", err)
		return err
	}
	facts, err := FetchCatFact()
	if err != nil {
		log.Printf("fetch error: %s\n", err)
		return err
	}

	for i, candle := range candles {
		fmt.Printf("%d:\n%s\n", i, candle.Text())
	}
	for i, fact := range facts {
		fmt.Printf("%d:\n%s\n", i, fact.Text())
	}
	return nil
}
