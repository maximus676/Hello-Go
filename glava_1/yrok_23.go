package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

func main() {
	u := User{}
	u.ID = 22
	u.Email = "test@test.com"
	u.FirstName = "John"

	bs, _ := json.Marshal(u)

	fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
	fmt.Println("============================")

	// ЗАдание  Реализуйте функцию Validate(req UserCreateRequest) string, которая валидирует запрос и возвращает строку с ошибкой "invalid request", если модель невалидна.
	//Если запрос корректный, то функция возвращает пустую строку. Правила валидации и структура модели описаны ниже. Не используйте готовые библиотеки и опишите правила самостоятельно.
	req := UserCreateRequest{
		FirstName: "Alice",
		Age:       25,
	}
	fmt.Println(Validate(req))

}

type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

func Validate(req UserCreateRequest) string { //UserCreateRequest тут как тип данных   req это переменная с моим объектом
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return "invalid request"
	} else if req.Age <= 0 || req.Age > 150 {
		return "invalid request2"
	}
	return ""
}
