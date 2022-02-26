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
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

// Función
func Saludar() string {
	return "Hola desde una función"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	//template, err := template.ParseFiles("index.html", "base.html")
	template, err := template.New("index.html").ParseFiles("index.html", "base.html")

	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"JavaScript"}
	cursos := []Curso{c1, c2, c3, c4}
	user := Usuario{"Pablo", 28, true, true, cursos}

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
