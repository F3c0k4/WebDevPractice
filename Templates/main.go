package main

import (
	"log"
	"os"
	"text/template"
)

var mytemp *template.Template

type dummy struct {
	A int
	B string
}

func init() {
	mytemp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	d := dummy{A: 4, B: "hello there"}
	e := []dummy{{A: 443, B: "I'm blue"}, {A: 321, B: "Da ba dee da ba die"}}
	mytemp.ExecuteTemplate(os.Stdout, "nodata.gohtml", nil)

	mytemp.ExecuteTemplate(os.Stdout, "datapassedsimply.gohtml", 22)

	mytemp.ExecuteTemplate(os.Stdout, "datathroughvar.gohtml", "this is my string")
	mytemp.ExecuteTemplate(os.Stdout, "dataloopedthrough.gohtml", []string{"the first", "the second", "the thrid", "and the fourth"})
	mytemp.ExecuteTemplate(os.Stdout, "dataloopedthroughvar.gohtml", []string{"the first", "the second", "the thrid", "and the fourth"})
	err := mytemp.ExecuteTemplate(os.Stdout, "datapassedthroughstruct.gohtml", d)
	if err != nil {
		log.Fatal(err)
	}
	err = mytemp.ExecuteTemplate(os.Stdout, "datapassedthroughstructlooped.gohtml", e)
	if err != nil {
		log.Fatal(err)
	}
}
