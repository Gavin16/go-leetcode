package main

import "fmt"

type BasicAnimal struct {
	name string
}

type Animal interface {
	Speak()
}

func (animal *BasicAnimal) Speak() {
	fmt.Println("animal", animal.name, "speak")
}

func (dog *Dog) Speak() {
	fmt.Println("dog speaks woof!")
}

type Dog struct {
	BasicAnimal
}

func main() {
	dog := Dog{BasicAnimal: BasicAnimal{name: "dog"}}
	dog.Speak()
	dog.BasicAnimal.Speak()
}
