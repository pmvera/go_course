package main

import "fmt"

// Relacion One2One
type User struct {
	name   string
	email  string
	status bool
}

type Student struct {
	user User
	code string
}

// Relacion One2Many
type Curse struct {
	title  string
	videos []Video
}

type Video struct {
	title string
	curse Curse
}

func main() {

	// Relacion One2One
	user1 := User{
		name:   "Pablo",
		email:  "pablo@gmail.com",
		status: true,
	}

	user2 := User{
		name:   "Pepe",
		email:  "pepe@gmail.com",
		status: true,
	}

	student1 := Student{
		user: user1,
		code: "0001",
	}

	fmt.Println(user2)
	fmt.Println(student1.user.name)

	// Relacion One2Many
	video1 := Video{title: "01-Introducción"}
	video2 := Video{title: "02-Instalación"}

	curso1 := Curse{
		title:  "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curse = curso1
	video2.curse = curso1

	fmt.Println(video1.curse.title)

	for _, video := range curso1.videos {
		fmt.Println(video.title)
	}
}
