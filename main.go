package main

import (
	"fmt"
	"math"
)

func main() {
	f := math.NaN()
	f2 := math.NaN()

	m := map[float64]string{
		f:  "Pedro",
		f2: "Pessoa",
	}

	fmt.Println(m)
	valor, ok := m[f]
	fmt.Println(valor, ok)
	delete(m, f)

	clear(m)
	fmt.Println(m)

	// m := map[string]string{
	// 	"Pedro":   "Pessoa",
	// 	"Joaquim": "Pedro",
	// }

	// clear(m)
	// fmt.Println(m)

	// m := make(map[string]string)
	// m["Pedro"] = "Pessoa"
	// valor, ok := m["Pedro"]
	// fmt.Println(valor, ok)

	// delete(m, "Pedro")
	// valor, ok = m["Pedro"]
	// fmt.Println(valor, ok)

	// m := map[string][]int{
	// 	"Pedro": {1, 2, 3},
	// }

	// m := map[string]string{
	// 	"Pedro":   "Pessoa",
	// 	"Joaquim": "Pedro",
	// }
	// m := make(map[string]string, 100)
	// fmt.Println(m)
}
