package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error){
	if err != nil {
		panic(err)
	}
}

func main(){
	// tempFile, err := os.CreateTemp("", "temporaryFile")
	// checkError(err)

	// fmt.Println("Temporary file created:", tempFile.Name())

	// defer tempFile.Close()
	// defer os.Remove(tempFile.Name())


	
	dir, err := os.Getwd()
	checkError(err)
	checkError(os.Mkdir("temp", 0755))
	dir = filepath.Join(dir, "/temp")

	tempDir, err := os.MkdirTemp(dir, "GoCourseTempDir-")
	checkError(err)
	for i := 1; i <= 2000; i++ {
		tempDir, err = os.MkdirTemp(dir, "GoCourseTempDir-")
		checkError(err)
		fmt.Println("Temporary directory created:", tempDir)
	}

	// defer os.RemoveAll(tempDir)


}
