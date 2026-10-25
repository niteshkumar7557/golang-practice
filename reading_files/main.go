package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// Open a file
	file, err := os.Open("output.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func (){
		fmt.Println("Closing file...")
		file.Close()
	}()

	fmt.Println("File opened successfully")


	// Read the contents of a opened file (as a whole)
	
	// data := make([]byte, 1024) // Buffer to read data into
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error while reading data from file:", err)
	// 	return
	// }
	// fmt.Println(string(data))



	// Read the contents of a opened file (line-by-line)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}


}