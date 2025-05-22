package main

import "fmt"

func main() {

	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println("===============")

	fmt.Println(fibbonachi(6))
	fmt.Println(fibbonachi(7))
	fmt.Println(fibbonachi(13))

}

func factorial(n int) int { // вычисление факторилала  4 * 3 * 2 * factorial(1)
	if n == 0 {
		return 1
	}
	return n * (factorial(n - 1))

}

func fibbonachi(f int) int {
	if f == 0 {
		return 0
	}

	if f == 1 {
		return 1
	}

	return fibbonachi(f-1) + fibbonachi(f-2) // как оно подставляются сюда чисда f от отдного до 13 ????? 
}
