package main

import "fmt"

func main() {
	var bilangan1, bilangan2, hasil int
	fmt.Scanln(&bilangan1, &bilangan2)
	for i := 0; i < bilangan2; i++ {
		if i == 0 {
			hasil = bilangan1
		} else {
			hasil = hasil * bilangan1
		}

	}
	fmt.Print(hasil)
}
