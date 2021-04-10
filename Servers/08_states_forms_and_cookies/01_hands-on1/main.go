package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", countVisitHandle)
	http.ListenAndServe(":8080", nil)
}

func countVisitHandle(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("default")

	if err == http.ErrNoCookie {
		c = &http.Cookie{Name: "default", Value: "0", Path: "/"}
	}

	counter, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatal(err)
	}
	counter++
	c.Value = strconv.Itoa(counter)
	http.SetCookie(w, c)
}
