package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Función
func Saludar(nombre string) string {
	return "Hola desde una función " + nombre
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	template, err := template.New("index.html").Funcs(funciones).ParseFiles("index.html")

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, nil)
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
