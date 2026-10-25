package main

import (
	"html/template"
	"os"
)

type Service struct {
	Name string
	Image string
	Ports []string
	Environment map[string]string
}

const tmplText = `services:
{{- range . }}
  {{.Name}}:
    image: "{{.Image}}"
    ports:
    {{- range .Ports }}
      - "{{.}}"
    {{- end }}
    environment:
    {{- range $key, $value := .Environment }}
      {{$key}}: "{{$value}}"
    {{- end }}
{{- end }}
`

func main(){

	services := []Service{
		{
			Name: "web",
			Image: "nginx:latest",
			Ports: []string{"80:80", "443:443"},
			Environment: map[string]string{
				"ENV": "production",
			},
		},
		{
			Name: "db",
			Image: "postgres-18",
			Ports: []string{"5432:5432"},
			Environment: map[string]string{
				"POSTGRES_USER": "postgres",
				"POSTGRES_DATABASE": "postgresdb",
				"POSTGRES_PASSWORD": "admin",
			},
		},
	}

	tmpl := template.Must(template.New("compose").Parse(tmplText))

	err := tmpl.Execute(os.Stdout, services)
	if err != nil {
		panic(err)
	}

}