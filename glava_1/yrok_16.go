package main

import "fmt"

func main() {

	var people = map[string]int{
		"Tom":   13,
		"Bob":   22,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people)        // map[Tom:1 Bob:2 Sam:4 Alice:8]
	fmt.Println(people["Bob"]) //22
	people["Bob"] = 33
	fmt.Println(people["Bob"]) //33

	fmt.Println("==================")
	if X, ok := people["awdawdvaw"]; ok { // ok — это булева переменная, которая будет равна true, если ключ "Tom" существует в map people // в X - запишется значение ключа Tom (13) // После инициализации переменных X и ok проверяется значение ok (которое на 3 месте стоит) и если оно тру пойдем условие fmt.Println(X)
		fmt.Println(X) // awdawdvaw такого ключа нет
	} else if X, ok := people["Bob"]; ok {
		fmt.Println(X) // Ключ Bob есть   // 33
	}

	fmt.Println("==================")
	//Для перебора элементов применяется цикл for:

	for key, value := range people {
		fmt.Println(key, value) //  Tom 13  Bob 33  Sam 4
	}

	people3 := make(map[string]int) //Функция make Она создает пустую хеш-таблицу:

	fmt.Println(people3) // map[]

	fmt.Println("==================222222")
	// Добавление и удаление элементов
	// Вариант 1  просто переназначить элемент
	people["Sam"] = 128 // изменит значение уже которое прописанов people
	people["Max"] = 222 // добавит новое значение в people
	fmt.Println(people) //map[Alice:8 Bob:33 Sam:128 Tom:13]

	// Вариант 2 для удаления применяется встроенная функция delete(map, key),
	delete(people, "Alice")
	fmt.Println(people) // map[Bob:33 Sam:128 Tom:13]

}
