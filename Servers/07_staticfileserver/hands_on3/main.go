package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("starting-files/public"))
	log.Fatal(http.ListenAndServe(":8080", fs))
}
