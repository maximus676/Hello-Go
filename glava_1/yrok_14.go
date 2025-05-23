package main

import "fmt"

func main() {

	defer fmt.Println(divide(15, 5)) // 3
	// defer fmt.Println(divide(15, 0)) // Y = 0 ошибка  //если данную строку раскомитать то проходка по  коду полностью остановится и до функций finish() вовсе не дойдёт

	finish()
	defer finish()  // позволяет выполнить определенную функцию в конце программы
	defer finish2() // Верхний defer выполнится последним так как при пробегании по пкоду встретился первым

	fmt.Println("случай 1")
	fmt.Println("случай 2")

}

func finish() {
	fmt.Println("случай 3")
}
func finish2() {
	fmt.Println("случай 4")
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Y = 0 ошибка")
	}

	return x / y
}

