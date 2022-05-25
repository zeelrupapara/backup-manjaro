package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func main() {
	db, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/information_schema")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println("Successfully connect to new database")
	
	db.AutoMigrate(&User{})
}
