package main

import (
	"fmt"
	"strings"
)

func statusByName(name string) string {
	// функция проверяет, что строка name начинается с подстроки "Mr."
	if strings.HasPrefix(name, "Mr.") {
		return "married man"
	} else if strings.HasPrefix(name, "Mrs.") {
		return "married woman"
	} else {
		return "single person"
	}
}

func main() {
	n := "Mr. Doe"
	fmt.Printf("%s is a %s\n", n, statusByName(n)) //
	// => Mr. Doe is a married man
	fmt.Println("===========================")

	n = "Mrs. Berry"
	fmt.Printf("%s is a %s\n", n, statusByName(n))
	// => Mrs. Berry is a married woman
	fmt.Println("===========================")

	n = "Karl"
	fmt.Printf("%s is a %s\n", n, statusByName(n))
	// => Karl is a single person
	fmt.Println("===========================")
	
	// Задание на if else
	fmt.Println(DomainForLocale("site.com", "ru"))  //ru
	fmt.Println(DomainForLocale("site.com", "SDA")) //by
	fmt.Println(DomainForLocale("site.com", "en"))  //en
	fmt.Println(DomainForLocale("site.com", ""))    // ge

}

func DomainForLocale(domain, locale string) string {
	if locale == "ru" {
		return "ru." + domain
	} else if locale == "en" {
		return "en." + domain
	} else if locale == "" { // проверка на приходящую пустую строку
		return "ge." + domain
	} else {
		return "by." + domain
	}

}
