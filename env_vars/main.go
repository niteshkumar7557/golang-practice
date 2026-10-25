package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("user:", user)
	fmt.Println("home:", home)

	os.Setenv("FRUIT", "apple")

	data := os.Environ()

	for _, i := range data{
		s := strings.Split(i, "=")
		fmt.Println(s[0])
	}


}
