package main

import (
	"log"

	"markettracker.com/tracker/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
