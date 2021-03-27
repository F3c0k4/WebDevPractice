package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	mytemp, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	mytemp.ExecuteTemplate(os.Stdout, "first.gohtml", nil)
}
