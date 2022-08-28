package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name":"f1",
			"last_name":"l1",
			"hair_color":"black",
			"has_dog":true
		},
		{
			"first_name":"f2",
			"last_name":"l2",
			"hair_color":"white",
			"has_dog":false
		}
	]`

	var unmarshalled []Person

	if err := json.Unmarshal([]byte(myJson), &unmarshalled); err != nil {
		log.Println("Unmarshaled failed", err)
	}

	log.Printf("data %v", unmarshalled)

	// write json from a struct

	var mySlice []Person = unmarshalled

	json, err := json.MarshalIndent(mySlice, "", "   ")

	if err != nil {
		log.Println("marshaled failed", err)

	}

	log.Printf("data %v", string(json))

}
