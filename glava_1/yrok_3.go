package main

import "fmt"

func main() {
	var a int = 4
	var b int = 3
	var c bool = a >= b
	var d bool = a != b
	fmt.Println(c) //true
	fmt.Println(d) //true

	var h bool = 4 > 5 && 6 > 8   //false
	var g bool = 3 <= 5 && 10 > 8 // true
	fmt.Println(h)
	fmt.Println(g)

	var i bool = 4 > 5 || 6 > 8   //false
	var u bool = 3 == 5 || 10 > 8 // true
	fmt.Println(i)
	fmt.Println(u)

	var f int = 2 << 2 // 10  на два разрядов влево = 1000 -в десятичной системе равно число 8.
	fmt.Println(f)     // 8

	var q int = 16 >> 3 // 16 в десятичной 10000 на три разряда вправо = 10 - что в десятичной системе представляет число 2
	fmt.Println(q)      //2
}
