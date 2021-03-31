package main

import (
	"fmt"
	"net/http"
)

type myhandler struct {
	str string
}

func main() {
	h := myhandler{
		"mystring",
	}
	http.ListenAndServe(":8080", h)
}

func (h myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Host+" "+h.str)
}
