package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer func(){
		fmt.Println("Closing file now...")
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	
	// keyword to filter lines
	keyword := "important"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword){
			fmt.Println("Filtered line:" , line)
			// updatedLine := strings.ReplaceAll(line, keyword, "***")
			// fmt.Println("Updated line:" , updatedLine)
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning the file:", err)
		return
	}


}