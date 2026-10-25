package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age"`
	Gender    string `json:"-"`
}

func main() {

	// person := Person{FirstName: "Nitesh", LastName: "Kumar", Age: 20}
	// person := Person{FirstName: "Nitesh", LastName: "", Age: 20}
	// person := Person{FirstName: "Nitesh", Age: 20}
	person := Person{FirstName: "Nitesh", LastName: "Kumar", Age: 20, Gender: "male"}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error while marshaling data")
	}

	fmt.Println(string(jsonData))

}
