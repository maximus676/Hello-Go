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

	return fibbonachi(f-1) + fibbonachi(f-2) // Проваливается функуию в вункуию пока функкия не раскладывается fibbonachi(1) или fibbonachi(0) после из условия выше сумирут еденицы или нули вверх по цепочке и в итоге получает число 5+3=8 (число фибаначи 6) 
}
