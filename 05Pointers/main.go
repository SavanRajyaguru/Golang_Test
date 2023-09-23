package main

import "fmt"

func main() {
	fmt.Println("pointer start here")

	num := 10
	fmt.Println("NUM: ", num)

	ptr := &num
	fmt.Println("PTR: ", ptr)
	fmt.Println("PTR: ", *ptr)

	newValue := *ptr + 2
	fmt.Println("NEW VALUE:", newValue)
}
