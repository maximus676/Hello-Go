package main

import (
	"fmt"
	"strconv" // функция позволяющая приобразовывать переменные в другие значения
)

// Пример приобразования входящего числа в строку
func main() {
	y := -54
	fmt.Println(IntToString(y))

	a, _ := strconv.Atoi("5.05")
	// a, _ := strconv.ParseFloat("5.05", 64)  // Пример с ParseFloat
	x := float64(a)
	fmt.Println(x) // выдаст не правильное значение и будет равен // 0 \\  так как Atoi Преобразует строку в целое число (int). нужно было тогда использовать ParseFloat

	// такая же проблема будет и с отрицательным числом
	g, _ := strconv.Atoi("-41")
	h := uint(g) // нужно использовать тип int и тогда ошибки не будет с отрицательным числом
	fmt.Println(h)
}

func IntToString(y int) string {
	s := strconv.Itoa(y)
	return s
}

// Преобразование строк в числа:
// Atoi(s string) (int, error): Преобразует строку в целое число (int).
// ParseInt(s string, base int, bitSize int) (i int64, err error): Преобразует строку в целое число с указанным основанием и размером.
// ParseFloat(s string, bitSize int) (float64, error): Преобразует строку в число с плавающей точкой (float64).

// Преобразование чисел в строки:
// Itoa(i int) string: Преобразует целое число (int) в строку.
// FormatInt(i int64, base int) string: Преобразует целое число в строку с указанным основанием.
// FormatFloat(f float64, fmt byte, prec, bitSize int) string: Преобразует число с плавающей точкой в строку с заданным форматом и точностью.

// Преобразование строк в булевы значения:
// ParseBool(str string) (bool, error): Преобразует строку в булево значение (bool).

// Преобразование булевых значений в строки:
// FormatBool(b bool) string: Преобразует булево значение в строку.
