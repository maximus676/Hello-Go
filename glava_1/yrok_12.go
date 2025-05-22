package main

import "fmt"

func main() {

	var f1 func(int, int) int = func(x int, y int) int { return x + y } // анонимная функция
	fmt.Println(f1(3, 4))                                               // 7
	fmt.Println("==========================================")

	action(10, 20, func(x, y int) int { return x + y })
	action(2, 5, func(x, y int) int { return x * y }) //Анонимная функция как аргумент функции

	fmt.Println("==========================================")
	var fun2 func(int, int) int = selectFn(2)
	fmt.Println(fun2(3, 3)) //9

	//Преимуществом анонимных функций является то, что они имеют доступ к окружению, в котором они определяются. Например:
	fmt.Println("==========================================")

	var fun3 func() int = square()

	fmt.Println(fun3()) // 9
	fmt.Println(fun3()) //16
	fmt.Println(fun3()) //25

}

func action(n1, n2 int, operation func(int, int) int) {
	var result int = operation(n1, n2)
	fmt.Println(result)

}

func selectFn(n1 int) func(int, int) int {
	if n1 == 1 {
		return func(x, y int) int { return x + y }
	} else if n1 == 2 {
		return func(x, y int) int { return x * y }
	} else {
		return func(x, y int) int { return x + y + 1000 }
	}
}

func square() func() int {
	var h int = 2 //3 // 4 // 5    // как сюда и по каокй причине перезаписывается переменная h ?????
	return func() int {
		h++
		return h * h
	}
}
