package main

import (
	"fmt"
	"strconv"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}

func main() {
	user := NewUser("Naoyoshi Aikawa", 29)
	fmt.Println(user)
}
