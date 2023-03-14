package main

import "fmt"

func main() {
	// funcao anonima
	func(x int) {
		fmt.Println(x)
	}(15)
}
