package main

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
)

func main() {
	s := "21"
	n := 21

	fmt.Println("string compare with int '21'==21 :", reflect.DeepEqual(s, n))

	slice1 := []int{1, 2, 1}
	slice2 := []int{1, 2, 1}
	fmt.Println("slice compare slice1==slice2 :", slices.Equal(slice1, slice2))

	map1 := map[int]int{1: 2, 2: 3, 3: 1}
	map2 := map[int]int{3: 1, 2: 3, 1: 2}

	fmt.Println("map compare map1==map2 :", maps.Equal(map1, map2))

	obj1 := struct {
		name  string
		addr  string
		score int
	}{"abc", "china", 100}

	obj2 := struct {
		name  string
		addr  string
		score int
	}{"abc", "china", 99}

	fmt.Println("struct compare obj1==obj2:", reflect.DeepEqual(obj1, obj2))

}
