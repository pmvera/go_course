package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Structs
type Usuario struct {
	UserName string
	Edad     int
}

var templates = template.Must(template.New("Template").ParseFiles("index.html", "base.html"))

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	//template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	user := Usuario{"Pablo", 28}
	err := templates.ExecuteTemplate(rw, "index.html", user)

	if err != nil {
		panic(err)
	}

}

func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}
	fmt.Println("Run server: http://localhost:9000/")
	// log.Fatal(http.ListenAndServe("localhost:9000", mux))
	log.Fatal(server.ListenAndServe())
}
