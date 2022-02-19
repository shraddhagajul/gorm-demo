package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})

	users := make([]User, 0)

	users = append(users, User{Username: "danny", FirstName: "Dan", LastName: "Morris"})
	users = append(users, User{Username: "manny", FirstName: "Manuel", LastName: "Chris"})
	users = append(users, User{Username: "john", FirstName: "John", LastName: "Doe"})

	for _, user := range users {
		db.Create(&user)
	}

	firstUser := User{}
	db.First(&firstUser)
	fmt.Println(firstUser)

	lastUser := User{}
	db.Last(&lastUser)
	fmt.Println(lastUser)

	fmt.Println("done")
}

type User struct {
	Id        uint
	Username  string
	FirstName string
	LastName  string
}
