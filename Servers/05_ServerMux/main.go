package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type myHandler struct {
	tplname string
}

func main() {
	//Exercise 1
	http.HandleFunc("/", handleSlash)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)

	//Exercise 2
	http.HandleFunc("/template", handleTemplate)

	//Exercise 3

	http.Handle("/template2", myHandler{"Data for Exercise 3"})
	http.ListenAndServe(":8080", nil)

}

//Exercise 1
func handleSlash(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "you requested just a /")
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "you requested /dog or something more")
}

func handleMe(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "you requested /me or something more")
}

//Exercise 2
func handleTemplate(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w, "tpl.gohtml", "this is the data passed in")
}

//Exercise 3
func (h myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "tpl.gohtml", h.tplname)
}
