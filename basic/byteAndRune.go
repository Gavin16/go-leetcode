package main

import "fmt"

func main() {
	s := "abcs"
	bytes := []byte(s)
	fmt.Println(bytes)

	runes := []byte(s)
	fmt.Println(runes)

	s1 := "你好world"
	bytes1 := []byte(s1)
	runes1 := []rune(s1)

	fmt.Println(bytes1)
	fmt.Println(runes1)
}
