package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
}

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

	// instead of ts.Execute, ExecuteTemplate to specifically want to respond using base template
	// struct is being used in nav.tmpl
	err = ts.ExecuteTemplate(w, "base", User{Name: "Ilia"})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
