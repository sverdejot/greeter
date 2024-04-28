package main

import (
	"log"

	"github.com/sverdejot/greeter/users/cmd/api/bootstrap"
)

func main() {
	errCh := bootstrap.Run()

	for err := range errCh {
		log.Printf("something happened: %v", err)
	}
}
