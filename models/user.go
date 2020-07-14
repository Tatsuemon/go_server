package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	user := User{
		1,
		"Naoyoshi Aikawa",
		29,
	}
	fmt.Println(user)
}
