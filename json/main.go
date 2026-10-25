package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email,omitempty"`
	Address Address `json:"address"`
}

type Employee struct {
	FullName string  `json:"full_name"`
	Emp_id   string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	person1 := Person{Name: "Nitesh", Age: 19}

	// Marshalling
	jsonData, err := json.Marshal(person1)
	if err != nil {
		log.Println("Error while marshalling")
		return
	}
	fmt.Println(string(jsonData))

	email := Person{Name: "Nitesh", Email: "nitesh@gmail.com"}
	jsonEmail, err := json.Marshal(email)
	if err != nil {
		log.Println("Error while marshalling")
		return
	}
	fmt.Println(string(jsonEmail))

	person2 := Person{Name: "Nitesh", Age: 19, Email: "nitesh@gmail.com", Address: Address{City: "Sonipat", State: "Haryana"}}
	jsonData, err = json.Marshal(person2)
	if err != nil {
		log.Println("Error while marshalling")
		return
	}
	fmt.Println(string(jsonData))

	jsonObj := `{"full_name": "Nitesh", "emp_id": "0009", "age": 20, "address": {"city": "Sonipat", "state": "Haryana"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonObj), &employeeFromJson)
	if err != nil {
		log.Println("Error while unmarshalling")
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println(employeeFromJson.Age)
	fmt.Println(employeeFromJson.Age + 5)
	fmt.Println(employeeFromJson.FullName)

	listOfCityState := []Address{
		{City: "Sonipat", State: "Haryana"},
		{City: "Rohini", State: "Delhi"},
		{City: "Jammu", State: "J and K"},
	}

	jsonData, err = json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error while unmarshalling")
	}

	fmt.Println(listOfCityState)
	fmt.Println(string(jsonData))


	jsonObj = `[
	{"full_name": "Nitesh", "emp_id": "0009", "age": 20, "address": {"city": "Sonipat", "state": "Haryana"}},
	{"full_name": "Rahul", "emp_id": "0010", "age": 21, "address": {"city": "Sonipat", "state": "Haryana"}}
	]`
	var employeesFromJson []Employee
	err = json.Unmarshal([]byte(jsonObj), &employeesFromJson)
	if err != nil {
		log.Println("Error while unmarshalling")
		return
	}
	fmt.Println(employeesFromJson)


	// Handeling unknown JSON structures
	jsonObj = `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]any

	err = json.Unmarshal([]byte(jsonObj), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}

	fmt.Println(data)
	fmt.Println(data["address"])
	address := data["address"].(map[string]any)
	fmt.Println(address)
	fmt.Println(address["city"])

}
