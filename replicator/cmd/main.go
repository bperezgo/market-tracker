package main

import (
	"log"

	"markettracker.com/replicator/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
