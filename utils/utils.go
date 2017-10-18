package utils

import (
	"log"
)

// FatalIfErr log.Fatal if err is different than nil
func FatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
