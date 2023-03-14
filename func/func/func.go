package main

import "fmt"

func main() {

	sum10 := funcReturnFunc(10)
	fmt.Println(sum10(20))

	fmt.Println(funcReturnFunc(10)(3))

}

func funcReturnFunc(y int) func(x int) int {
	return func(x int) int {
		return x + y
	}
}
