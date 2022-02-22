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
  db.Create(&user)

	// db.Debug().Delete(&user)
	db.Debug().Model(&user).Where("first_name = ? ", "Joe").Delete(&user)

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
