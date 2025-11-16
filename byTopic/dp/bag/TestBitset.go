package main

import (
	"fmt"
	"math/big"
)

func main() {
	one := big.NewInt(7)
	fmt.Println(one.BitLen())
	p := big.NewInt(1)
	fmt.Println(p.Lsh(one, 3))
	fmt.Println(p)
	f := big.NewInt(1)
	f.Or(f, p.Lsh(one, 3))
	fmt.Println(f)
}
