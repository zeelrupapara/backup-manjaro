
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func BenchmarkOrmCreate(b *testing.B) {
	db, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/information_schema")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	users := stubUsers(b)
	for _, user := range users {
		db.Create(user)
	}
}
func genRandomString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func stubUsers(b *testing.B) (users []*User) {
	for i := 0; i < b.N; i++ {
		user := &User{
			Name:     genRandomString(100),
			Password: genRandomString(100),
		}
		users = append(users, user)
	}

	return users
}
