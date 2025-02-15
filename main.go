package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	foo(&arr)
	fmt.Println(arr)
}

func foo(slice *[4]int) {
	slice[0] = 123

	// _ = slice[3]
	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
	// fmt.Println(slice[2])
	// fmt.Println(slice[3])

}
