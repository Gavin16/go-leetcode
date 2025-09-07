package main

func main() {

	balanceCh := make(chan int)
	resultCh := make(chan int)

	go func() {
		balance := 0
		for {
			select {
			case amount := <-balanceCh:
				balance += amount
			case resultCh <- balance:
			}
		}
	}()

	for i := 0; i < 10; i++ {
		balanceCh <- 100
	}

	println("最后余额:", <-resultCh)
}
