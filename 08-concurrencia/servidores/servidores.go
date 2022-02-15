package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(servidor string, channel chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		//fmt.Println(servidor, "no está disponible")
		channel <- servidor + "no está disponible"
	} else {
		//fmt.Println(servidor, "está disponible")
		channel <- servidor + "está disponible"
	}
}

func main() {
	inicio := time.Now()

	channel := make(chan string)

	servidores := []string{
		"https://oregoom.com",
		"https://www.google.com",
		"https://www.udemy.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
	}

	for _, servidor := range servidores {
		go checkServer(servidor, channel)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-channel)
	}

	tiempo_paso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución:", tiempo_paso)
}
