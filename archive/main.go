package main

import (
	"fmt"
)

type User struct {
	id   string
	name string
	age  int
}

func (u *User) IncrementAge() {
	u.age++
}

func main() {
	user := &User{
		id:   "1",
		name: "John Lennon",
		age:  30,
	}

	user.IncrementAge()
	fmt.Println(user.age)
}
