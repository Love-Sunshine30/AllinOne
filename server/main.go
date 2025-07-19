package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed internal/*
var content embed.FS

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/download/", download)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)

}
