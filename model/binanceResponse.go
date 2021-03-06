package model

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// BinanceCandle is a struct for minimal needed data of one candle
type BinanceCandle struct {
	OpenTime int64
	High     int
	Low      int
}

// BinanceCandlesResponse is a struct for json unmarshalling
type BinanceCandlesResponse struct {
	BinanceCandles []BinanceCandle
}

// UnmarshalJSON is a custom unmarshall function for one array -> candle
func (c *BinanceCandle) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Printf("Error while decoding %v\n", err)
		return err
	}
	var err error = nil
	c.OpenTime = int64(v[0].(float64))
	if i, _ := strconv.ParseFloat(v[2].(string), 64); err != nil {
		log.Printf("Error while parsing %v\n", err)
		return err
	} else {
		c.High = int(i)
	}
	if i, _ := strconv.ParseFloat(v[3].(string), 64); err != nil {
		log.Printf("Error while parsing %v\n", err)
		return err
	} else {
		c.Low = int(i)
	}

	return err
}

// UnmarshalJSON is a custom unmarshall function for the whole query json
func (c *BinanceCandlesResponse) UnmarshalJSON(data []byte) error {
	var v []BinanceCandle
	if err := json.Unmarshal(data, &v); err != nil {
		log.Printf("Error whilde decoding %v\n", err)
		return err
	}
	c.BinanceCandles = v
	return nil
}

// Text for printing out one candle
func (c BinanceCandle) Text() string {
	p := fmt.Sprintf("CloseTime: %d\nHigh: %d\nLow: %d\n", c.OpenTime, c.High, c.Low)
	return p
}
