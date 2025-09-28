package main

import "sort"

type MyInt int64

type Heap struct {
	sort.IntSlice
	MyInt
}

func main() {
	h := &Heap{}
	println(h.MyInt)
	h.MyInt = 1
	println(h.MyInt)
}
