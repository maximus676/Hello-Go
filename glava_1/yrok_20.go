package main

import (
	"fmt"
	"strings"
)

func main() {

	flag := true
	text := "hello"

	fmt.Println(flag || text != "")

	if "text" == "text2" {
		fmt.Println("++++++++++")
	}
	fmt.Println("===========================")
	// задачи из урока на условия

	IsValid(0, "hello world")   // false
	IsValid(-22, "hello world") // false
	IsValid(22, "")             // false
	IsValid(225, "hello world") // true
	fmt.Println("===========================")
	//Строки
	//Строки Для слияния строки с переменной нужно использовать fmt.Sprintf():  и для передачи переменнолй используется структура %s",
	username := "Ivan"
	fmt.Println(fmt.Sprintf("hello,,, %s", username)) // "hello, Ivan"
	fmt.Println(len(username))                        // ответ 4 - это количество битов
	fmt.Println("===========================")

	name := " ивАн             "

	fmt.Println(Greetings(name))

}

func IsValid(id int, text string) bool {
	if id > 0 && text != "" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	return true
}

func Greetings(name string) string {
	greetings := "Привет, "
	name2 := strings.Trim(name, " ")
	name2 = strings.ToLower(name2)
	name2 = strings.Title(name2)
	return greetings + name2 + "!"
}

//Привет, Иван!
