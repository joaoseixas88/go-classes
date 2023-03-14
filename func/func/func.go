package main

import "fmt"

func main() {

	sum := funcReturnFunc(10)
	fmt.Println(sum(20))

}

func funcReturnFunc(y int) func(x int) int {
	return func(x int) int {
		return x + y
	}
}
