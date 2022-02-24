package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)
type Name struct {
	FirstName string   //field name should be similar to User struct in order to use scan
	LastName string
	Calendars string `gorm:"column:name"`
}
func main() {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})
	db.DropTable(&User{})
	db.CreateTable(&User{})


	users := make([]User, 0)

	users = append(users, User{Username: "danny", FirstName: "Dan", LastName: "Morris"})
	users = append(users, User{Username: "manny", FirstName: "Manuel", LastName: "Chris"})
	users = append(users, User{Username: "john", FirstName: "John", LastName: "Doe"})

	for _, user := range users {
		db.Create(&user)
	}

	calendars := make([]Calendar, 0)
	calendars = append(calendars, Calendar{UserID: 1, Name: "tour1"})
	calendars = append(calendars, Calendar{UserID: 2, Name: "tour2"})
	calendars = append(calendars, Calendar{UserID: 3, Name: "tour3"})
	for _, calendar := range calendars {
		db.Create(&calendar)
	}

	usersInDb := make([]User, 0)
	db.Debug().Model(&User{}).Limit(2).Find(&usersInDb)
	//For pagination
	db.Debug().Model(&User{}).Limit(2).Offset(1).Find(&usersInDb)
	//inflate specific fields 
	//If we use select, other value apart for "first_name" & "last_name" will be nil or 0 
	db.Debug().Model(&User{}).Select([]string{"first_name", "last_name"}).Find(&usersInDb)
	fmt.Println(&usersInDb)
	//inflate specific column
	usernames := []string{}
	db.Debug().Model(&User{}).Pluck("username", &usernames)
	fmt.Println(usernames)
	names := []Name{}

	//inflate multiple columns. Only fetches given columns
	db.Debug().Model(&User{}).Select([]string{"first_name","last_name"}).Scan(&names)
	fmt.Println(names)
	//Count
	countUsersInDb := 0
	db.Debug().Model(&User{}).Count(&countUsersInDb)
	fmt.Println(countUsersInDb)

	//inflate default values if query fails 
	defaultUser := User{}
	db.Debug().Model(&User{}).Where("username = ?", "sammuel").Attrs(&User{FirstName: "Elton"}).FirstOrInit(&defaultUser)
	fmt.Println("Default user : ",defaultUser)

	//To override : scope is in application only, not db
	db.Debug().Model(&User{}).Where("username = ?", "manny").Assign(&User{FirstName: "Elton"}).FirstOrInit(&defaultUser)
	fmt.Println("Default user : ",defaultUser)

	userCalDet := []Name{}
	db.Debug().Model(&User{}).Joins("inner join calendars on calendars.user_id = users.id").Select("users.first_name, users.last_name, calendars.name").Scan(&userCalDet)
}

type User struct {
	Id        uint
	Username  string
	FirstName string
	LastName  string
	Calendar Calendar
}

type Calendar struct {
	UserID uint
	Name string
}
