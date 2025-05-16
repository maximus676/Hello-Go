package main

import "fmt"

func main() {
	var c int = 3
	var v int = 4

	var a = add(c, v)  // 7
	var b = add(20, 6) // 26
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("============================")
	var r = add2(30, 20)
	fmt.Println(r) // 50
	fmt.Println("============================")
	var age, name = add3(20, 25, "Max", "lev")
	fmt.Println(age)
	fmt.Println(name)

	//тоже самое но функция переписана более развернуто с переменными 
	var age1, name1 = add4(20, 25, "Max", "lev")
	fmt.Println(age1)
	fmt.Println(name1)
}

func add(x, y int) int { // второй int это уже от return "тип_возвращаемого_значения "
	return x + y
}

func add2(x, y int) (z int) { // можно указать сразу что я хочу вернуть в тип_возвращаемого_значения и не прописывать в return и написать в конце просто return
	z = x + y
	return // не ошибка написать и так "return z"
}

func add3(x, y int, firstName, lastName string) (z int, full string) {
	z = x + y
	full = firstName + " " + lastName
	return z, full
}

func add4(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var full string = firstName + " " + lastName
	return z, full
}
