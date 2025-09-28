package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"Name" db:"name"`
	Email string `json:"email" db:"email"`
}

func main() {
	user := User{Name: "hello", Email: "abc@gmail.com"}
	ut := reflect.TypeOf(user)
	field, _ := ut.FieldByName("Name")
	fmt.Println(field.Tag.Get("json"))
	fmt.Println(field.Tag.Get("db"))
}
