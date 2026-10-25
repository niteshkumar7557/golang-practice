package main

import (
	"html/template"
	"os"
)

func main(){
	// tmpl, err := template.New("example").Parse("Hi {{.name}}\n")
	// if err != nil{
	// 	panic("failed at parsing")
	// }

	tmpl := template.Must(template.New("ex").Parse("Hello, {{.nane}}!\n"))

	data := map[string]any{
		"name": "Nitesh",
	}
	err := tmpl.Execute(os.Stdout, data)
	if err != nil{
		panic("failed at execution")
	}
}