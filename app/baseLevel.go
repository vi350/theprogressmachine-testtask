package app

import (
	"fmt"
	"log"
)

func BaseLevel() {
	candles, err := FetchCrypto("USDT", "BTC")
	if err != nil {
		log.Println(err)
	}
	facts, err := FetchCatFact()
	if err != nil {
		log.Println(err)
	}

	for i, candle := range candles {
		fmt.Printf("%d:\n%s\n", i, candle.Text())
	}
	for i, fact := range facts {
		fmt.Printf("%d:\n%s\n", i, fact.Text())
	}
}
