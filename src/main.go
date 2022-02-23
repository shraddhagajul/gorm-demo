package main

import (

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
	users := []User{}
	// db.Debug().Where("username = ?", "manny").Find(&users)
	// db.Debug().Where(&User{Username: "danny"}).Find(&users)
	// db.Debug().Where(map[string]interface{}{"username":"manny"}).Find(&users)

	// db.Debug().Where("username in (?)", []string{"manny", "danny"}).Find(&users)
	// db.Debug().Where("username like (?)", "%man%").Find(&users)

	// db.Debug().Not("username = ?", "manny").Find(&users)
	db.Debug().Where("username = ?", "manny").Or("username = ?", "danny").Find(&users)

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