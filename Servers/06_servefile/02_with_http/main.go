package main

import (
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/honda", serveHonda)
	http.ListenAndServe(":8080", nil)
}

func serveHonda(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("honda.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found", 404)
	}

	//http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
	http.ServeFile(w, req, fi.Name())
}
