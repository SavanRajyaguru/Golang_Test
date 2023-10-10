package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex

var urls = []string{
	"https://www.google.com",
	"https://instagram.com",
	"https://youtube.com",
	"https://x.com",
	"https://meet.google.com",
	"https://meet.ai.com",
	"https://example.com",
	"https://designs.ai",
}

func makeHttpRequest(url string, i int) {
	http.Get(url)
	fmt.Println(url)
	defer wg.Done()
}

var counter = 0
var Slice = make([]int, 0)

func increaseCounter(i int) {
	defer wg.Done()
	// fmt.Println(i)
	mut.Lock()
	// mut.TryLock()
	Slice = append(Slice, i)
	mut.Unlock()
	counter += 1
}

func main() {
	// for i, url := range urls {
	// 	wg.Add(1)
	// 	go makeHttpRequest(url, i)
	// }

	// the counter does not go to 100;
	// it should go to 100 right?
	// for i := 1; i <= 10; i += 1 {
	// 	wg.Add(1)
	// 	go increaseCounter(i)
	// }
	// for i := -10; i <= 0; i += 1 {
	// 	wg.Add(1)
	// 	go increaseCounter(i)
	// }
	// time.Sleep(5 * time.Second)
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
		c1 <- "222"
		c1 <- "111"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case _, ok := <-c1:
			if ok {
				for _, val := range <-c1 {
					fmt.Println("received", val)
				}
			}
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// fmt.Println(Slice)

	// fmt.Println(counter + 1)
}
func worker(input chan int, output chan int, done chan bool) {
	defer wg.Done()
	for {
		select {
		case n := <-input:
			// Do some work on n
			output <- n * 2
		case <-done:
			close(output)
			return
		}
	}
}
