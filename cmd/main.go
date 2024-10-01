package main

import (
	"log"

	"./LiveDb"
)

func main() {
	// mapping Users info (mock data, real will come from JSON request)
	User := map[string]string{
		"name": "Ni3",
		"age":  "19",
	}

	DB, err := LiveDb.NewLiveDB()
	if err != nil {
		log.Fatal(err)
	}

	collection, err := DB.createCollection("users")
	if err != nil {
		log.Fatal(err)
	}
}
