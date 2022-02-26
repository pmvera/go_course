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
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handler Error
func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Render template
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		handlerError(rw, http.StatusInternalServerError)
	}
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	user := Usuario{"Pablo", 28}
	renderTemplate(rw, "index.html", user)
}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)

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
