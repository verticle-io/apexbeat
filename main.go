package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/verticle-io/apexbeat/beater"
)

func main() {
	err := beat.Run("apexbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
