package main

import "fmt"

func main() {
	var a int8 = 5
	var b int8 = 8
	if a > b {
		fmt.Println("NO")
	} else if a == b {
		fmt.Println("====")
	} else {
		fmt.Println("YES")
	}

	var c int = 3
	switch {

	case c == 7:
		fmt.Println("c = 7")
	case c == 8:
		fmt.Println("c = 8")
	case c == 6, c == 5, c == 3:
		fmt.Println("c = 6 или 5 или 4, но это не точно")
		fallthrough
	case c < 1:
		fmt.Println("c < 1")
		fallthrough
	case c < 2:
		fmt.Println("c < 2")
		fallthrough // тут уже не совсем корректно ставиить так как выполниться последующие default что быдет не верно
	default:
		fmt.Println("значение переменной c не определено")
	}

}
