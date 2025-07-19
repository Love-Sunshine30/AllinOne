package main

import(
	"html/template"
	"net/http"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("ui/index.html"));
	tmpl.Execute(w, nil);
}

func main(){
	fs := http.FileServer(http.Dir("./internal"));
	http.Handle("/internal/", http.StripPrefix("/internal/", fs));

	http.HandleFunc("/", home);
	
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil);

}