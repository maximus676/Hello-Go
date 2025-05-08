package main

import "fmt"

func main() {
	var hello string = "start"
	hello = "Hello world"
	text := "Max"
	var tab string = "\tТом \nСой\\ер"

	var age uint8
	fmt.Println(tab)
	fmt.Println(hello)
	fmt.Println(len(text))

	var name string
	fmt.Println("wfat is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + " !")
	fmt.Println("How old are you ?")
	fmt.Scan(&age)
	fmt.Println("you are " + fmt.Sprint(age) + " yers !") //число со строкой выводить нельзя по этому используем Sprint

}
