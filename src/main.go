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

	pdb := db.DB()
	defer pdb.Close()
	err = pdb.Ping()
	if err != nil {
		panic(err.Error())
	}
	
	fmt.Println("Connection to databse connection")
}

