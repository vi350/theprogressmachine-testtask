package main

import (
	"github.com/vi350/theprogressmachine-testtask/app"
	"log"
)

func main() {
	if err := app.BaseLevel(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
