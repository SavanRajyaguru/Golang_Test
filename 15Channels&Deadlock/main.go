package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	// 1 is the buffer value means sender can send only one extra value of the receiver
	myCha := make(chan int, 2)

	// RECEIVES ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		// fmt.Println("Go1: ", &ch)

		val := <-ch
		fmt.Println(val)
		// fmt.Println(isTrue)
		val1 := <-ch
		fmt.Println(val1)
		// fmt.Println(isTrue2)
		fmt.Println(<-ch)
		fmt.Println(<-ch)

	}(myCha, wg)

	// SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer close(ch)
		defer wg.Done()
		// fmt.Println("Go2: ", &ch)

		ch <- 5
		ch <- 6
		ch <- 7
		// ch <- 8
		// ch <- 7
	}(myCha, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer close(ch)
		defer wg.Done()
		ch <- 50
		// ch <- 60
		// fmt.Println(&ch)
		// ch <- 7
		// ch <- 8
		// ch <- 7
	}(myCha, wg)

	// func() {
	// 	defer fmt.Println("Done")
	// 	defer fmt.Println("close")
	// }()
	// fmt.Println(<-myCha)
	// fmt.Println("Final: ", &myCha)
	ch := make(chan int, 10)
	go fib(10, ch)
	for x := range ch {
		fmt.Println(">>", x)
	}
	wg.Wait()

}

func fib(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
