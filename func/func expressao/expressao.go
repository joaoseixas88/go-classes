package main

import "fmt"

func main() {
	// funcao como expressao
	function := func(x int) {
		fmt.Println(x)
	}

	function(10)
}
