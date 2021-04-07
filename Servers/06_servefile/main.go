package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/honda", handleHonda)

	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "something something")
}

func handleHonda(w http.ResponseWriter, req *http.Request) {
	pic, err := os.Open("honda.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer pic.Close()

	io.Copy(w, pic)
}
