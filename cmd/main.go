package main

import (
	"fmt"
	"log"

	"github.com/Ni-011/LiveDb/LiveDb"
)

func main() {
	// mapping Users info (mock data, real will come from JSON request)
	User := map[string]string{
		"name": "Ni3",
		"age":  "19",
	}

	_ = User

	DB, err := LiveDb.NewLiveDB()
	if err != nil {
		log.Fatal(err)
	}

	collection, err := DB.CreateCollection("users")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(collection)
}
