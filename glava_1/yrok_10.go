package main

import "fmt"

func isEven(n int) bool { // функции с одинаковыми типами для пеердачи их в функцию sum как  criteria
	return n%2 == 0
}
func isPositive(n int) bool { // функции с одинаковыми типами для пеердачи их в функцию sum как  criteria
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {

	result := 0
	for i := 0; i < len(numbers); i++ {
		if criteria(numbers[i]) {
			result += numbers[i]
		}
	}

	// for _, val := range numbers{     // пример цикла без индекса 
    //     if(criteria(val)){
    //         result += val
    //     }
    // }

	return result

}

func main() {

	var slice []int = []int{2, 6, 3, -1, 7, -4, 23}

	var sumOfEvens int = sum(slice, isEven) // сумма четных чисел   //sumOfEvens int  - потому что функция sum возвращает int
	fmt.Println(sumOfEvens)                 // 4

	sumOfPositives := sum(slice, isPositive) // сумма положительных чисел
	fmt.Println(sumOfPositives)              // 37
}
