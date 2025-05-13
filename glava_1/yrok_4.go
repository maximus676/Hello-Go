package main

import "fmt"

func main() {
	var test [3]int = [3]int{1, 2, 3}
	fmt.Println(test) // [1 2 3]
	var a int8 = 3
	var b int8 = 3

	var isAlive2 bool = a == b // true

	var numbers [3]int = [3]int{1, 2, 3}
	var numbers2 [3]int = [3]int{1, 2, 3}
	var isAlive bool = numbers == numbers2 // true

	// var numbers3 [3]int = [3]int{1, 2, 3}
	// var numbers4 [3]int = [3]int{1, 2, 4}
	// var isAlive3 bool = numbers3 == numbers4 // ошибка Длинна масива является частью его типа  (int) по этому два массив с разным колчичесвтом "число_элементов" сравнивать нельзя

	fmt.Println(isAlive)  // true
	fmt.Println(isAlive2) // true
	// fmt.Println(isAlive3) // ошибка Длинна масива является частью его типа  (int) по этому два массив с разным колчичесвтом "число_элементов" сравнивать нельзя

	var r [4]int = [4]int{1, 1, 2, 4}
	// r = [5]int{1, 4, 5, 6, 6}  так же и с приссвоением если массиву пытаться присвоить более длинный массив у них будут разные типы и упадёт ошибка.
	r = [4]int{1, 4, 5, 6}
	fmt.Println(r)

	var names [3]string = [3]string{"max", "olga", "polle"}
	fmt.Println(names)
	fmt.Println(names[2])
	names[1] = "Eva"
	fmt.Println(names[1])

	var indeks [3]string = [3]string{2: "PPP", 0: "DDD", 1: "MMM"}
	fmt.Println(indeks[2]) //PPP

}
