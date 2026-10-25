package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main(){
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"


	// Join paths using filepath.Join
	joinedPath := filepath.Join("home", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized path:", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	fmt.Println("Base:", filepath.Base("/home/user/docs/file.txt"))
	
	fmt.Println(filepath.IsAbs(relativePath))
	fmt.Println(filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	
	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(absPath)

}