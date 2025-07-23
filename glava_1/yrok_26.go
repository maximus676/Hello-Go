package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// важно инициализировать слайс той же длины
	numsCp := make([]int, len(nums))
	// numsCp := []int{}  // ТАк будет ошибочно и вернет пустой массив не скланировав массив nums
	// numsCp := make([]int, 3) // скопирует новый массив с 3 элементам от исходника

	copy(numsCp, nums)
	fmt.Println(numsCp) // [1,2,3,4,5]
	fmt.Println("======================")

	// ЗАДАНИЕ
	//Реализуйте функцию IntsCopy(src []int, maxLen int) []int, которая создает копию слайса src с длиной maxLen. Если maxLen равен нулю или отрицательный,
	//то функция возвращает пустой слайс []int{}. Если maxLen больше длины src, то возвращается полная копия src.

	src := []int{1, 2, 3, 4}
	maxLen := 7
	fmt.Println(IntsCopy(src, maxLen))

}

func IntsCopy(src []int, maxLen int) []int {

	if maxLen <= 0 {
		return []int{}
	} else if maxLen > len(src) {
		newSrc2 := make([]int, len(src))
		copy(newSrc2, src)
		return newSrc2
	} else {  // Если в деапощзоне длинны массива src то отдай кусок соответствующий числу maxLen
		newSrc := make([]int, maxLen)
		copy(newSrc, src)
		return newSrc
	}

}
