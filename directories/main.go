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

	checkError(os.Mkdir("subdir", 0755))
	// defer os.RemoveAll("subdir")

	checkError(os.WriteFile("subdir/file.txt", []byte(""), 0755 ))

	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))
	checkError(os.MkdirAll("subdir/parent/child4", 0755))

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result{
		fmt.Println(entry.Name(), entry.IsDir())
	}

	checkError(os.Chdir("subdir/parent/child1"))

	result, err = os.ReadDir(".")
	checkError(err)

	for _, entry := range result{
		fmt.Println(entry.Name(), entry.IsDir())
	}

	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)




	// filepath.Walk and filepath.WalkDir

	pathfile := "subdir/"

	fmt.Println("Walking Directory")

	checkError(filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		checkError(err)
		fmt.Println(path)
		return nil
	}))


	checkError(os.RemoveAll("subdir"))

}
