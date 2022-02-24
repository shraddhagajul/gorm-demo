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
	seedDb(db)

	users := []User{}

	db.Debug().Preload("Calendar.Appointments").Find(&users)

	for _, u := range users {
		fmt.Printf("\n%v\n", u.Calendar)
	}



}

func seedDb(db *gorm.DB) {
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
		Username:  "joe_jonas",
		FirstName: "Joe",
		LastName:  "Jonas",
		Calendar: Calendar{
			Name: "2022 itinerary",
			Apointments: []Appointment{
				{Subject: "Tour New York", Attendees: users},
				{Subject: "Tour Washington", Attendees: users},
			},
		},
	})
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
	Name        string
	UserID      uint
	Apointments []Appointment `gorm:"polymorphic:owner"`
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	OwnerID     uint
	OwnerType   string
	Attendees   []User `gorm:"many2many:appointment_user"`
}

type TaskList struct {
	gorm.Model
	Appointments []Appointment `gorm:"polymorphic:owner"`
}
