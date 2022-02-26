package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//db.Ping()
	fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	//db.TruncateTable("users")
	//user := models.CreateUser("Moreno", "moreno1234", "moreno@gmail.com")
	//fmt.Println(user)

	users := models.ListUsers()
	fmt.Println(users)
	user := models.GetUser(2)
	fmt.Println(user)

	user.Username = "Pepe"
	user.Save()
	users = models.ListUsers()
	fmt.Println(users)

	user.Delete()
	users = models.ListUsers()
	fmt.Println(users)

	db.Close()
}
