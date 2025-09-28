package main

func main() {
	m := map[int]int{}
	m[1], m[2] = 11, 22
	m[3] = 0

	v, ok := m[2]
	println("value=", v)
	if ok {
		delete(m, 2)
		println("m[2] has been deleted.")
		println("m[2]=", m[2])
	} else {
		println("m[2] not exists! ")
	}

}
