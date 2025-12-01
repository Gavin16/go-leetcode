package main

import "fmt"

//func main() {
//	one := big.NewInt(7)
//	fmt.Println(one.BitLen())
//	p := big.NewInt(1)
//	fmt.Println(p.Lsh(one, 3))
//	fmt.Println(p)
//	f := big.NewInt(1)
//	f.Or(f, p.Lsh(one, 3))
//	fmt.Println(f)
//
//	fmt.Println("---------strconv --------")
//	//n, err := strconv.Atoi("0012")
//	n, err := strconv.Atoi("")
//	if err != nil {
//		fmt.Printf("%v", err)
//	} else {
//		fmt.Println(n)
//	}
//}

func main() {

	a := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	for _, r := range a {
		r[0], r[1], r[2] = r[0]+1, r[1]+1, r[2]+1
	}
	fmt.Println(a)
}
