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
	// db.SingularTable(true)
	db.CreateTable(&User{})

	db.CreateTable(&Engineer{})

	for _, f := range db.NewScope(&Engineer{}).Fields() {
		fmt.Println(f.Name)
	}

	db.Model(&User{}).AddIndex("idx_first_name", "first_name")
	db.Model(&User{}).AddUniqueIndex("idx_last_name", "LastName")

	db.Model(&User{}).RemoveIndex("idx_first_name")

}

type User struct {
	// gorm.Model
	UserId    int    `gorm:"primary_key"`
	Username  string `sql:"type:VARCHAR(15);unique;unique_index; not null"`
	FirstName string `sql:"size:150;DEFAULT:'Nobody'"`
	LastName  string `gorm:"column:LastName"`
	Count     int    `gorm:"AUTO_INCREMENT"`
	TempField bool   `sql:"-"`
}

// func (u User) TableName() string {
// 	return "stakeholders"
// }

type Engineer struct {
	UserDetails User `gorm:"embedded"`
	Degree string
}
