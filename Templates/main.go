package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var mytemp *template.Template

type dummy struct {
	A int
	B string
}

var fm = template.FuncMap{
	"TU":       strings.ToUpper,
	"fdateMDY": monthDayYear,
	"times4":   times4,
	"sub3":     sub3,
}

func monthDayYear(t time.Time) string {
	return t.Format("2006-01-02")
}

func times4(num int) int {
	return num * 4
}

func sub3(num int) int {
	return num - 3
}

func init() {

	mytemp = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
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

	err = mytemp.ExecuteTemplate(os.Stdout, "functionpassedin.gohtml", "this_should_turn_to_upper_case")
	if err != nil {
		log.Fatal(err)
	}

	err = mytemp.ExecuteTemplate(os.Stdout, "formattime.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	err = mytemp.ExecuteTemplate(os.Stdout, "pipeline.gohtml", 4)
	if err != nil {
		log.Fatal(err)
	}

	err = mytemp.ExecuteTemplate(os.Stdout, "defined_template_used.gohtml", "which received the data I sent")
	if err != nil {
		log.Fatal(err)
	}

}
