package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./internal/ui"))
	http.Handle("/internal/ui/", http.StripPrefix("/internal/ui/", fs))

	http.HandleFunc("/", home)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)

}
