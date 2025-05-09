package main

import "fmt"

func main() {
	var a int = 10
	var b int = 4
	var c int = a / b
	fmt.Println(c) // 2 покажет только целое число (хоть ответ 2.5)

	var k float32 = 10
	var l float32 = 4
	var m float32 = k / l
	fmt.Println(m) // 2 покажет только целое число (хоть ответ 2.5)

	var d float32 = 11 / 4   // 2  результат деления будет округляться до целого числа, даже если результат присваивается переменной типа float32/float64:
	var e float32 = 11 / 4.0 // 2.75  Решение данной проблемы
	fmt.Println(d)           // 2   результат деления будет округляться до целого числа, даже если результат присваивается переменной типа float32/float64:
	fmt.Println(e)           // 2.75

	a++ //Постфиксный инкремент (x++) как и в JS
	fmt.Println(a) //11
	a--   //10
	a--   //9
	fmt.Println(a)
}
