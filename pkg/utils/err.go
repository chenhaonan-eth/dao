package utils

import (
	"log"
)

func CatchError(detail string, err error) {
	if err != nil {
		log.Printf("%s: %v\n", detail, err)
	}
}
