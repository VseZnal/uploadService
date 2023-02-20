package errors

import (
	"log"
)

func LogError(err error) error {
	log.Println(err)

	return err
}

func HandleFatalError(err error, message string) {
	log.Println(err)

	log.Fatal(message)
}
