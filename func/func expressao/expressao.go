package main

import "fmt"

func main() {
	// funcao como expressao => igual a const func do js
	function := func(x int) {
		fmt.Println(x)
	}

	function(10)
}
