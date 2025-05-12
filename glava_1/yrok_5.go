package main

import (
	"fmt"
)

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
	switch c + 2 {
	case 7:
		fmt.Println("c = 7")
	case 8:
		fmt.Println("c = 8")
	case 6, 5, 4:
		fmt.Println("c = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной c не определено")
	}

}
