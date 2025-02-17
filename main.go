package main

import (
	"errors"
	"fmt"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

var ErrNotFound = errors.New("Not found")

// func RaizQuadrada(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, SqrtError{"Não existe raiz quadrada de número negativo!"}
// 	}
// 	resultado := math.Sqrt(x)
// 	return resultado, nil
// }

func main() {

	err := foo()
	var sqtrError *SqrtError
	if err != nil && errors.As(err, &sqtrError) {
		fmt.Println(sqtrError.msg)
		return
	}
	fmt.Println("caiu fora")

	// err := foo()
	// if err != nil && errors.Is(err, ErrNotFound) {
	// 	fmt.Println("deu erro not found")
	// 	return
	// }
	// fmt.Println("foi para fora")

	// x := -10
	// res, err := RaizQuadrada(float64(x))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)
}

func foo() error { return &SqrtError{msg: "teste"} }
