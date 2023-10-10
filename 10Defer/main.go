package main

import "fmt"

func main() {
	fmt.Println("====== Defer Start ======")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	fmt.Println("Hello")
	myDefer()
}

// Three, Two, One (Stack LIFO)
// Hello, myDefer(), Three, Two, One
// myDefer: 4, 3, 2, 1, 0

func myDefer() {
	i := 0
	for i < 5 {
		defer fmt.Println(i)
		i++
	}
}
