package main

import (
	"fmt"
	"time"
)

//func main() {
//	balanceCh := make(chan int)
//	resultCh := make(chan int)
//	stopCh := make(chan struct{})
//	go func() {
//		balance := 0
//		for {
//			select {
//			case amount := <-balanceCh:
//				balance += amount
//			case resultCh <- balance:
//			case <-stopCh:
//				fmt.Println("worker exited")
//				return
//			}
//		}
//	}()
//	for i := 0; i < 10; i++ {
//		balanceCh <- 100
//	}
//	println("最后余额:", <-resultCh)
//	close(stopCh)
//}

//func main() {
//	dataCh := make(chan int)
//	stopCh := make(chan int)
//	go func() {
//		for {
//			select {
//			case v := <-dataCh:
//				fmt.Println("收到数据: ", v)
//			case <-time.After(1500 * time.Millisecond):
//				fmt.Println("超时未收到数据")
//				return
//			case v, ok := <-stopCh:
//				fmt.Printf("停止执行 v= %v, ok = %v\n", v, ok)
//				return
//			}
//		}
//	}()
//	for i := range 10 {
//		time.Sleep(1 * time.Second)
//		dataCh <- i + 1
//	}
//	close(stopCh)
//	time.Sleep(3 * time.Second)
//}

func main() {
	var chA = make(chan int)
	var chB = make(chan int)
	var chC = make(chan int)

	asyncCallA := func() {
		time.Sleep(800 * time.Millisecond)
		chA <- 10
	}
	asyncCallB := func() {
		time.Sleep(500 * time.Millisecond)
		chB <- 20
	}
	asyncCallC := func() {
		time.Sleep(1300 * time.Millisecond)
		chC <- 30
	}
	go asyncCallA()
	go asyncCallB()
	go asyncCallC()

	timeout := time.After(2 * time.Second)
	for i := 0; i < 3; i++ {
		select {
		case r := <-chA:
			fmt.Printf("chanA return result = %v\n", r)
			chA = nil
		case r := <-chB:
			fmt.Printf("chanB return result = %v\n", r)
			chB = nil
		case r := <-chC:
			fmt.Printf("chanC return result =%v\n", r)
			chC = nil
		case <-timeout:
			fmt.Println("超时未完成调用, 退出执行")
			return
		}
	}
}
