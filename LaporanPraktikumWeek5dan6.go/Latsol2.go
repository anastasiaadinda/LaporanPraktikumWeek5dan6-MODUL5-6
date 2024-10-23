package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var V, r, t float64
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&r, &t)

		V = (1.0 / 3.0) * (math.Pi * r * r * t)
		fmt.Println(V)
	}
}
