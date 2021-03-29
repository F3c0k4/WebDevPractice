package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name, Address, City, Zip, Region string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main() {
	hotels := []Hotel{
		{Name: "My hotel_1", Address: "Somewhere St 12", City: "San Francisco", Zip: "28432", Region: "Southern"},
		{Name: "My hotel_2", Address: "Another St 12", City: "Los Angeles", Zip: "22342", Region: "Central"},
	}

	f, err := os.Create("result.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = tmpl.Execute(f, hotels)
	if err != nil {
		log.Fatal(err)
	}
}
