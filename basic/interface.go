package main

import "fmt"

type Eat interface {
	taste()
}
type Smell interface {
	smell()
}

type Enjoy interface {
	Eat
	Smell
}

type Fruit struct {
}

type Apple struct {
}

type Peach struct {
}

func (apple Apple) taste() {
	fmt.Println("taste sweet..")
}

func (peach Peach) taste() {
	fmt.Println("taste scent..")
}

func testInterSwitch(i Eat) {
	switch v := i.(type) {
	case Apple:
		fmt.Println("this is an apple")
	case Peach:
		fmt.Println("this is a peach")
	default:
		fmt.Println("unknow type: %T\n", v)
	}
}

func main() {
	type User struct {
		name   string
		nation string
	}
	var i interface{} = User{
		name:   "hh",
		nation: "CN",
	}

	s, ok := i.(User)
	if ok {
		fmt.Println(s)
	}

	var a Eat
	a = Apple{}
	testInterSwitch(a)

	a = Peach{}
	testInterSwitch(a)
}
