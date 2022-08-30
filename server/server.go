package server

import (
	"fmt"
	"log"
	"net/http"
	"flag"
)

func Server() {
	mux := http.NewServeMux()
	
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
    flag.Parse()
    listener := gateway.ListenAndServe
    portStr := "n/a"
    if *port != -1 {
        portStr = fmt.Sprintf(":%d", *port)
        listener = http.ListenAndServe
        http.Handle("/", http.FileServer(http.Dir("./ui")))
    }


	fileServer := http.FileServer(http.Dir("./ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art-web", result)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/authors", authors)

// 	fmt.Println("Starting the web server .... on http://localhost:5050")
// 	err := http.ListenAndServe(":5050", mux)
// 	log.Fatal(err)
}
