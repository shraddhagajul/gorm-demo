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

	db.DropTable(&User{})
	// db.SingularTable(true)
	db.CreateTable(&User{})

}

type User struct {
	Id        uint
	Username  string
	FirstName string
	LastName  string
}

func (u User) TableName() string {
	return "stakeholders"
}
