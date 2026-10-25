package main

import (
	"embed"
	"fmt"
)

//go:embed example.txt
var content string

//go:embed data_folder
var dataFolder embed.FS

func checkError(err error){
	if err != nil {
		panic (err)
	}
}

func main(){

	fmt.Println("Embeded content:", content)

	content, err := dataFolder.ReadFile("data_folder/hello.txt")
	checkError(err)

	fmt.Println("Embeded file content:", string(content))

}
