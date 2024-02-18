package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type FontMap struct {
	Name string
}

func createFileUsingTemplate(t *template.Template, filename string, data interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var tmplFile = "template.html"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	var data []FontMap

	entries, err := os.ReadDir("./font-maps/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range entries {
		mapName := strings.Replace(file.Name(), ".map", "", 1)
		if mapName == "ML-TTKarthika" {
			continue
		}
		data = append(data, FontMap{Name: mapName})
	}

	createFileUsingTemplate(tmpl, "index.html", data)
}
