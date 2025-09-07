package main

import (
	"time"
	"unicode/utf8"
)

func ticker(ch chan int) {
	i := 0
	for {
		ch <- i
		i++
		time.Sleep(500 * time.Millisecond)
	}
}

func RuneCount(b []byte) int {
	count := 0
	for i := 0; i < len(b); {
		if b[i] < utf8.RuneSelf {
			i++
		} else {
			_, n := utf8.DecodeRune(b[i:])
			i += n
		}
		count++
	}
	return count
}

func main() {
	s := "你好gopher"
	r := []byte(s)
	println("number of rune:", RuneCount(r))
}

//func basic() {
//	ch := make(chan int)
//	stop := make(chan struct{})
//
//	go func() {
//		ticker(ch)
//	}()
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		close(stop) // 3秒后停止循环
//	}()
//
//	for {
//		select {
//		case v := <-ch:
//			fmt.Println(v)
//		case <-stop:
//			fmt.Println("Stopped")
//			fmt.Println(utf8.RuneSelf)
//			return
//		}
//	}
//}
