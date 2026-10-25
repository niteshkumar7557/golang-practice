package main

import (
	"flag"
	"fmt"
	"os"
)



func main(){

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args{
		fmt.Printf("Argument %v: %v\n", i, arg)
	}

	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Nitesh", "Name of the user")
	flag.IntVar(&age, "age", 20, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male?:", male)

}
