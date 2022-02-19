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

	user := User{
		Username: "test_user",
		FirstName: "Test",
		LastName: "User",
	}

	db.Create(&user)

	fmt.Println("done")
	fmt.Println(user)
}

type User struct {
	Id uint
	Username string
	FirstName string
	LastName string
}
