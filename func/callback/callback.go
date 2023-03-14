package main

import "fmt"

func main() {
	x := func(number int) int {
		return number * 5
	}
	callBack(x, true)
	callBack(x, false)
}

func callBack(x func(y int) int, condition bool) {
	if condition {
		fmt.Println("Condition is true")
		fmt.Println(x(35))
	} else {
		fmt.Println("Condition is False")
	}
}
