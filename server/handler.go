package main

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(content, "internal/ui/index.html"))
	tmpl.Execute(w, nil)
}

func download(w http.ResponseWriter, r *http.Request) {
	//get the pdf name
	pdfName := r.URL.Path[len("/download/"):]
	// read the file
	data, err := content.ReadFile("internal/pdfs/" + pdfName)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set headers to force download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+pdfName+"\"")
	w.Write(data)
}
