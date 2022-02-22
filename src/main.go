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
	db.CreateTable(&User{})

	user := &User{Username: "joe_jonas", FirstName: "Joe", LastName: "Jonas"}
 
	tx := db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
	}

	user.LastName = "Mcmillan"
	if err := tx.Save(&user).Error; err == nil {
		tx.Rollback()
	}

	tx.Commit()

}

type User struct {
	gorm.Model           
	Username     string
	FirstName    string
	LastName     string

}

// type Appointment struct {
// 	gorm.Model
// 	UserID      uint
// 	Subject     string
// 	Description string
// 	StartTime   time.Time
// 	Length      uint
// }
