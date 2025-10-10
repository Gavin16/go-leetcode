package main

import (
	"fmt"
	"reflect"
)

type TestType struct {
	slice []int
	dict  map[string]int
	num   int
}

func main() {
	obj1 := new(TestType)
	println(obj1.slice == nil)
	obj1.slice = append(obj1.slice, 221)
	for _, n := range obj1.slice {
		println(n)
	}
	println(obj1.dict == nil)
	var dict map[string]int
	dict2 := new(map[string]int)

	fmt.Println(reflect.TypeOf(dict))
	fmt.Println(reflect.TypeOf(dict2))
	fmt.Println(reflect.TypeOf(obj1.dict))

	nums := []int{1, 2, 3}
	fmt.Println(nums)
}
