package main

import (
	"errors"
	"fmt"
)

func main() {

	user, err := NewUser(true)
	if err != nil {
		fmt.Println("Algum erro na hora de crirar o User")
		return
	}
	user.Foo()

	// a := 10
	// b := 0
	// res, err := dividir(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(a, "/", b, res)

}

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("Um erro")
	}
	return &User{}, nil
}

// func dividir(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("Nao pode dividir por zero")
// 	}
// 	return a / b, nil

// }
