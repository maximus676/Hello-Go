package main

import "fmt"

func main() {
	hello()
	plus(2, 5)
	plus(3, 7)
	var v int = 30
	fmt.Println("v before: ", v)
	increment(v)
	fmt.Println("v after: ", v) //значение переменной V никак не изменилось. Потому что при вызове функции передается копия значения переменной.
	add(1, 2, 3)                // sum = 6
	add(1, 2, 3, 4)             // sum = 10
	var lor [5]int = [5]int{5, 6, 7, 2, 3}
	var lor2 [4]int = [4]int{1, 2, 3, 4}
	var lor3 []int = []int{5, 6, 7, 2, 3} // Если нет жёского количества динны массива то пеерводить в срез не нужно  "[:]"
	add2(lor[:])                          // sum = 15  Таким образом можно передать массив нужно указывать "[:]" приобразоывая  массив в срез так акак у массива может быть разная длинна и эторазные типы
	add2(lor2[:])                         // sum = 10  lor2[:]
	add2(lor3)                            // sum =  23   // не нужен  "[:]" так как нет жёского типа с длинной массива []int
	add2([]int{5, 6, 5, 7, 3})            // sum = 26
}

func hello() {
	fmt.Println("Hello World")
}

func plus(a int, b int) {
	var c = a + b
	fmt.Println("a + b = ", c)
}

func increment(x int) {
	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}
func add(props ...int) {
	var sum = 0
	for i := 0; i < len(props); i++ {
		sum += props[i]
	}
	fmt.Println(" ============== sum = ", sum)
}

func add2(props []int) {
	var sum = 0
	for i := 0; i < len(props); i++ {
		sum += props[i]
	}
	fmt.Println(" add2 ============== sum2  = ", sum)
}
