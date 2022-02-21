package main

import (
	"fmt"
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
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	db.Debug().Model(&Calendar{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	users := []User{
		{Username: "manny"},
		{Username: "danny"},
		{Username: "nick"},
	}

	for i := range users {
		db.Save(&users[i])
	}

	db.Debug().Save(&User{
		Username: "joe_jonas",
		FirstName: "Joe",
		LastName: "Jonas",
		Calendar: Calendar{
			Name: "2022 itinerary",
			Apointments: []Appointment{
				{Subject: "Tour New York", Attendees: users },
				{Subject: "Tour Washington", Attendees: users},
			},
		},
	})

	

	u:= &User{}
	c:= &Calendar{}
	db.First(&u).Related(&c, "calender")
	//Note : while using first...related : gorm doesnt inflate the child object by default 
	fmt.Println(u)
	// &{{1 2022-02-21 07:39:49.377786 +0000 UTC 2022-02-21 07:39:49.377786 +0000 UTC <nil>} joe_jonas Joe Jonas {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>}  0}}
	fmt.Println()
	fmt.Println(c)
}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name   string
	UserID uint
	Apointments []Appointment
}

type Appointment struct {
	gorm.Model
	Subject string
	Description string
	StartTime time.Time
	Length uint
	CalendarID uint
	Attendees []User    `gorm:"many2many:appointment_user"`
}
