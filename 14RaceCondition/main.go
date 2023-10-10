package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}
var mut = &sync.RWMutex{} // good way to use RWMutex in code compare to just Mutex

func main() {
	score := make([]int, 0)

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Four")
		mut.RLock()
		fmt.Println(score)
		// time.Sleep(1 * time.Second)
		mut.RUnlock()
		defer wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("Data:", score)

}
