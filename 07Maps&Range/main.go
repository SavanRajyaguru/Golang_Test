package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("====== Maps starts =====")

	m := make(map[string]int)
	m["One"] = 1
	m["Two"] = 2
	m["Three"] = 3

	fmt.Println("Map:", m)
	fmt.Println("Length:", len(m))

	delete(m, "One")
	fmt.Println("Map:", m)

	clear(m)
	fmt.Println("Map:", m)

	rs, err := m["One"]

	fmt.Println("Data:", rs)
	fmt.Println("Data:", err)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
