package main

import "fmt"

func main() {
	var bilangan int
	var hasil int = 0
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		hasil += i
	}
	fmt.Println(hasil)
}
