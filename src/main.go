package main

import (
	"time"

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
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	user := &User{Username: "joe_jonas", FirstName: "Joe", LastName: "Jonas", Salary: 1000}
	user1 := &User{Username: "dylan_23", FirstName: "Dylan", LastName: "Lockhart", Salary: 2000}
	appointments := []Appointment{
		{Subject: "Concert Tour"},
		{Subject: "Drive Tour"},
	}
	user.Appointments = appointments

	db.Debug().Create(&user)
	db.Debug().Create(&user1)

	// user.FirstName = "Shelly"
	// user.LastName = "Gardner"

	// db.Debug().Save(&user)

	// db.Debug().Model(&user).Update("first_name", "Shelly")

	// db.Debug().Model(&user).Updates(map[string]interface{}{
	// 	"first_name": "Shelly",
	// 	"last_name": "Gardner",
	// })

	db.Debug().Table("users").Where("salary < 1500").Update("salary", gorm.Expr("salary + 200"))

}

type User struct {
	Id           uint
	Username     string
	FirstName    string
	LastName     string
	Salary uint
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	UserID      uint
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
}
