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

var templates = template.Must(template.New("Template").ParseGlob("templates/**/*.html"))

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	user := Usuario{"Pablo", 28}
	err := templates.ExecuteTemplate(rw, "index.html", user)

	if err != nil {
		panic(err)
	}

}
func Registro(rw http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(rw, "registro.html", nil)

	if err != nil {
		panic(err)
	}

}

func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}
	fmt.Println("Run server: http://localhost:9000/")
	// log.Fatal(http.ListenAndServe("localhost:9000", mux))
	log.Fatal(server.ListenAndServe())
}
