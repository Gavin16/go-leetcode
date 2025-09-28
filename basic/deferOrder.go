package main

import "fmt"

type T struct {
	Name string `json:"name" db:"name"`
}

func (t *T) pp() {
	fmt.Println(t.Name)
}

func main() {
	//testDefer(1)

	ts := []T{{"a"}, {"b"}, {"c"}}
	for _, v := range ts {
		defer v.pp()

	}
}

func testDefer(a int) {
	defer fmt.Println("1, a=", a)                    // 在defer语句注册时立即求a的值
	defer func(v int) { fmt.Println("2, a=", v) }(a) // 在defer语句注册时立即求a的值
	defer func() { fmt.Println("3, a=", a) }()       // 注册时只使用了a的引用,实际等到执行才会获得a的最新值
	a++
}
