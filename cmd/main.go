package main

import (
	"fmt"
	"log"

	"github.com/Ni-011/LiveDb/LiveDb"
)

func main() {
	// mapping Users info (real will come from JSON request)
	User := map[string]string{
		"name": "Ni3",
		"age":  "19",
	}

	// initate LiveDb isntance
	DB, err := LiveDb.NewLiveDB()
	if err != nil {
		log.Fatal(err)
	}

	// creates and inserts user data into a collection
	id, err := DB.Insert("Users", User)
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Printf("%+v", id)
}
