package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice == nil)

	// slice := []int{1, 2, 3, 4, 5}
	// slice2 := slice[:0]
	// fmt.Println(slice2, len(slice2), cap(slice2))

	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:4]
	// arr[2] = 15
	// slice[0] = 123
	// fmt.Println(arr)
}
