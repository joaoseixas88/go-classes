package main

import "fmt"

func main() {
	recursive(0)
	fmt.Println("fatorial", fatorial(45))

}

func recursive(x int) int {
	if x >= 50 {
		return x
	}
	x++
	fmt.Println(x)

	return recursive(x)
}

func fatorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return x * fatorial(x-1)
}
