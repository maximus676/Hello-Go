package main

import (
	"errors"
	"fmt"
)

// func main() {

// 	x := 2

// 	// ПРИМЕР С Ошибкой
// 	fmt.Println("Какое ваше число?")
// 	fmt.Scan(&x)
// 	fmt.Println(checkPositive(x))

// }


func main() {
	// Пример использования функции checkPositive
	number, err := checkPositive(-5)   // две переменных 1 число возращаемое 2 ошибка или ее отсутвие 
	if err != nil {
		// Обрабатываем ошибку
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Число является положительным:", number)
	}

	// Пример с положительным числом
	number, err = checkPositive(10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Число является положительным:", number)
	}
}

func checkPositive(number int) (int, error) {
	if number < 0 {
		// Возвращаем ошибку, если число отрицательное
		return 0, errors.New("число должно быть положительным")
	}
	// Возвращаем число, если оно положительное
	return number, nil // nil, указывая на отсутствие ошибки
}