package main

import "fmt"

func main() {

	users := []string{"Tom", "Alice", "Kate"}

	fmt.Println(users[2])

	for _, value := range users {
		fmt.Println(value)

	}

	fmt.Println("==============================")
	users2 := make([]string, 3) //С помощью функции make() можно создать срез из нескольких элементов, которые будут иметь значения по умолчанию:
	users2[0] = "Max"
	users2[1] = "Toood"
	users2[2] = "Bob"

	fmt.Println(users2)
	fmt.Println("==============================")
	//Добавление элемента в срез происходит при помощи функции append    append(slice, value)
	users3 := []string{"Tom1", "Alice1", "Kate1"}
	users3 = append(users3, "Bob1")

	for _, value := range users3 {
		fmt.Println(value)
	}

	fmt.Println("==============================")
	//Оператор среза s[i:j]  с i по (j-1  не включительно )

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

	users4 := initialUsers[0:3]
	users5 := initialUsers[7:8]

	fmt.Println(users4) // [Bob Alice Kate]
	fmt.Println(users5) //[Robert]
}
