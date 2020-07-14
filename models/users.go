package models

import (
	"fmt"
	"strconv"

	"github.com/Tatsuemon/go_server/db"
)

// User defines an user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

// NewUser creates user instance
func NewUser(name string, age int) *User {
	user := &User{
		Name: name,
		Age:  age,
	}
	db.Get().Create(user)
	fmt.Println(*user)
	return user
}

func GetUser(id int) *User {
	var user *User
	db.Get().Where("ID = ?", 1).Firsst(&user)
	return user
}

func AllUsers() []*User {
	var users []*User
	db.Get().Find(&users)
	return users
}

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}
