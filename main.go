package main

import (
	"github.com/discobean/go-docker-containerid/containerid"
	"log"
)

func main() {
	id, err := containerid.Find()

	if err != nil {
		log.Fatalf("Failed to get container ID: %v\n", err)
	}

	log.Printf("Got Container ID: %s\n", id)
}

