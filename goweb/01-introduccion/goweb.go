package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El método es: " + r.Method)
	fmt.Fprintln(rw, "Hola mundo")
}

func NotResponsePage(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "La página no funciona", http.StatusNotFound)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola, %s tu edad es %s", name, age)
}

func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", NotResponsePage)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// Router
	// http.HandleFunc("/", Hola)
	// http.HandleFunc("/page", NotResponsePage)
	// http.HandleFunc("/error", Error)
	// http.HandleFunc("/saludar", Saludar)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}
	fmt.Println("El servidor está ejecutándose en el puerto 9000")
	fmt.Println("Run server: http://localhost:9000/")
	// log.Fatal(http.ListenAndServe("localhost:9000", mux))
	log.Fatal(server.ListenAndServe())
}
