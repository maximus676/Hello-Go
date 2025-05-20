package main

import "fmt"

func add(x int, y int) int      { return x + y }
func minus(x int, y int) int    { return x - y }
func multiply(x int, y int) int { return x * y }

func operation(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return minus
	} else {
		return multiply
	}

}

func main() {

	var different_functions func(int, int) int = operation(1)
	fmt.Println(different_functions(2, 4))   //6

	different_functions = operation(2)
	fmt.Println(different_functions(10, 1))  //9

	different_functions = operation(6)
	fmt.Println(different_functions(5, 5))  //25

}
