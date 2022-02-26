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

// Función
func Saludar() string {
	return "Hola desde una función"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	//template, err := template.ParseFiles("index.html", "base.html")
	template, err := template.New("index.html").ParseFiles("index.html", "base.html")

	user := Usuario{"Pablo", 28}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
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
	fmt.Println("El servidor está ejecutándose en el puerto 9000")
	fmt.Println("Run server: http://localhost:9000/")
	// log.Fatal(http.ListenAndServe("localhost:9000", mux))
	log.Fatal(server.ListenAndServe())
}
