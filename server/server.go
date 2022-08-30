package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art-web", result)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/authors", authors)

	fmt.Println("Starting the web server .... on http://localhost:5050")
	err := http.ListenAndServe(":5050", mux)
	log.Fatal(err)
}
