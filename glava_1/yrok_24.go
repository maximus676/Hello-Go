package main

import "fmt"

func main() {

	nums := [5]int{1, 2, 3, 4, 5}
	i := 4
	val := 23
	fmt.Println(SafeWrite(nums, i, val))

	// Теория СЛАЙСЫ
	fmt.Println("======================")
	// есть нюанс с изменением массива (слайса) в функции из за привышения вместимости слайса может ссылаться на другой уже автоматом созданный слайс
	nums3 := []int{1, 2, 3, 4, 5} // вместимость 5

	modifySlice(nums3)

	fmt.Println(nums3) // [1 2 10 4 5]
	fmt.Println("====================== обход")
	//Как проблему с вместимостью обойти сделать элемент с запасом вместимости
	nums4 := make([]int, 0, 6)
	nums4 = append(nums4, 1, 2, 3, 4, 5) // вместимость 5
	fmt.Println("Вместимовсть:", cap(nums4))
	modifySlice(nums4)

	fmt.Println(nums4) // [1 2 10 4 5]

	//задание 2 СЛАЙСЫ
	fmt.Println("====================== ЗАДАНИЕ 2")
	slice := []int{}
	fmt.Println(RemoveFirstElement(slice))
	slice = []int{2, 1, 4, 5}
	fmt.Println(RemoveFirstElement(slice))
}

//Реализуйте функцию SafeWrite(nums [5]int, i, val int) [5]int, которая записывает значение val в массив nums по индексу i,
// если индекс находится в рамках массива. В противном случае массив возвращается без изменения.
//ЗАДАНИЕ 1
func SafeWrite(nums [5]int, i, val int) [5]int {

	if i < 5 && i > 0 {
		nums[i] = val
		return nums
	} else {
		return nums
	}

}

//Теория
func modifySlice(nums3 []int) {
	nums3[2] = 10            // элемент изменится и в исходном слайсе
	nums3 = append(nums3, 6) // элемент не добавится в исходный слайс, так как превысили изначальную вместимость и nums теперь ссылается на новый массив
	nums3[3] = 15            // элемент НЕ изменится в исходном слайсе
}

// Теория ^^^

//ЗАДАНИЕ 2
//Реализуйте функцию RemoveFirstElement(slice []int) []int, которая удаляет первый элемент из слайса. Если в функцию передан пустой слайс, то функция также вернет пустой слайс.
// original := []int{1, 2, 3, 4, 5}
// RemoveFirstElement(original) // [2 3 4 5]
// RemoveFirstElement([]int{}) // []

func RemoveFirstElement(slice []int) []int {
	if len(slice) <= 0 {
		return slice
	} else {
		return slice[1:]
	}
}
