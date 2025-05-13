package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i * i)
	}
	fmt.Println("==========================")
	//Можно убрать изменение счетчика в само тело цикла и оставить только условие:
	var x int = 1
	for x < 4 {
		fmt.Println(x + x)
		x++
	}
	fmt.Println("==========================")
	// Циклы могут быть вложенными, то есть располагаться внутри других циклов. Например, выведем на консоль таблицу умножения:

	for r := 1; r <= 10; r++ {
		for j := 1; j < 10; j++ {
			fmt.Print(r*j, "\t")
		}
		fmt.Println()
	}
	fmt.Println("==========================")
	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}
	// Если не нужны индексы ставим значению index нижний прочерк _
	// var users = [3]string{"Tom", "Alice", "Kate"}
	// for _, value := range users {
	// 	fmt.Println(, value)
	// }

	fmt.Println("==========================")
	// Так же альтенативныфй вариант
	var names [3]string = [3]string{"MAX", "ALAX", "HENK"}
	for g := 0; g < len(names); g++ {
		fmt.Println(fmt.Sprint(g) + " " + names[g])
	}

	// Операторы break и continue
	fmt.Println("+++++++++++++++++++++++")
	// Например нам нужно сложить только положительные числа (continue)
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum int = 0
	for t := 0; t < len(numbers); t++ {
		if numbers[t] < 0 {
			continue
		}
		sum += numbers[t]
	}
	fmt.Println("Sum: ", sum) //27

	// Пример если число масива больше 10 закакончить работу цикла. (break)
	var numbers2 = [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2}
	var sum2 = 0

	for p := 0; p < len(numbers2); p++ {
		if numbers2[p] > 4 {
			break //  если бы использовал continue то зацепил бы последнюю цифру 2 в конце масива так не обрывал работу массива и был бы результат 12
		}
		sum2 += numbers2[p]
	}
	fmt.Println("Sum2:", sum2) // 10
}
