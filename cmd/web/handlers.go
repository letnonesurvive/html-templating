package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl", // must be the first one
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...) // the
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil) // instead of ts.Execute, ExecuteTemplate to specifically want to respond using base template
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
