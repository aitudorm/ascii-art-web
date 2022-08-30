package server

import (
	"ascii-art-web/ascii"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/templates/homePageN.html", "./ui/templates/resultPageN.html"}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
	err = tmpl.ExecuteTemplate(w, "homePageN.html", nil)
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
}

func result(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-web" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("ui/templates/resultPageN.html")
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
	textInput := r.FormValue("inputData")
	bannerType := r.FormValue("banner")
	data, errorCheck := ascii.AsciiLogic(textInput, bannerType)

	if errorCheck {
		http.Error(w, "Bad Request, 400", http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
		return
	}
	err = tmpl.ExecuteTemplate(w, "resultPageN.html", data)
	if err != nil {

		http.Error(w, "Internal Error", 500)
		return
	}
}

func authors(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/authors" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("ui/templates/authors.html")
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
		return
	}
	err = tmpl.ExecuteTemplate(w, "authors.html", nil)
	if err != nil {

		http.Error(w, "Internal Error", 500)
		return
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("ui/templates/about.html")
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
		return
	}
	err = tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {

		http.Error(w, "Internal Error", 500)
		return
	}
}
