package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func display(message string) {
	fmt.Println(message)
}

func action(x int, y int, operation func(int, int) int) {  // функция которая принемает функции похожего типра которыми можно пользоваться из выхова функции action
	var result = operation(x, y)
	fmt.Println(result)
}

func main() {

	f := add             //или так var f func(int, int) int = add
	fmt.Println(f(3, 4)) // 7

	f = multiply         // теперь переменная f указывает на функцию multiply
	fmt.Println(f(3, 4)) // 12

	// f = display      // ошибка, так как функция display имеет тип func(string)

	var f1 func(string) = display // норм
	f1("hello")

	fmt.Println("========================")

	action(10, 15, add)
	action(10, 10, multiply)

}
