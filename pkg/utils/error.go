package utils

import (
	"fmt"
	"log"
)

func CheckError(err error, context string) {
	if err != nil {
		log.Fatalf("Error %s: %v", context, err)
	}
}

func NewError(message string) error {
	return fmt.Errorf(message)
}
