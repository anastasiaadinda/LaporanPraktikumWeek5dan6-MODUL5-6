package main

import "fmt"

func main() {
	var bilangan int
	var hasil int = 1
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		hasil = hasil * i

	}
	fmt.Println(hasil)
}
