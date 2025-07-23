package main

import "fmt"

func main() {
	strs := []string{"Max", "Polle", "Goha"}

	fmt.Println(Map(strs, mapFunc)) // [MaxPPP PollePPP GohaPPP]
	fmt.Println(strs)               // [Max Polle Goha]

}

// Реализуйте функцию func Map(strs []string, mapFunc func(s string) string) []string, которая преобразует каждый элемент слайса
//  strs с помощью функции mapFunc и возвращает новый слайс. Учтите, что исходный слайс, который передается как strs, не должен измениться в процессе выполнения.

func Map(strs []string, mapFunc func(s string) string) []string {
	Newstrs := []string{}
	for i := 0; i < len(strs); i++ {
		Newstrs = append(Newstrs, mapFunc(strs[i]))
	}
	return Newstrs
}

func mapFunc(s string) string {
	return s + "PPP"
}


