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
	seedDb(db)

	u := User{}
	// db.Debug().First(&u)	
	// db.Debug().FirstOrInit(&u, &User{Username: "inTheHouse"})
	// fmt.Println(u)
	db.Debug().FirstOrCreate(&u, &User{Username: "inTheHouse"})
	fmt.Println(&u)

}

type User struct {
	Id        uint
	Username  string
	FirstName string
	LastName  string
}

func seedDb(db *gorm.DB) {
	db.DropTable(&User{})
	db.CreateTable(&User{})

	users := make([]User, 0)

	users = append(users, User{Username: "danny", FirstName: "Dan", LastName: "Morris"})
	users = append(users, User{Username: "manny", FirstName: "Manuel", LastName: "Chris"})
	users = append(users, User{Username: "john", FirstName: "John", LastName: "Doe"})

	for _, user := range users {
		db.Create(&user)
	}
}