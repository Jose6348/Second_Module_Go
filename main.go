package main

import "fmt"

type User struct {
	name string
	ID   uint64
}

func (u *User) UpdateName(newName string) {
	u.name = newName
}

func UpdateName(u *User, newName string) {
	u.name = newName
}

func main() {
	f := foo.Foo()
	user := User{name: "Pedro", ID: 10}
	user.UpdateName("Joaquim")
	UpdateName(&user, "Pedro")
	fmt.Println(user)
}
