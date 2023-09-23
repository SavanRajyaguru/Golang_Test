package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("===== Arrays starts =====")

	var arr [5]int
	fmt.Println("Arr", arr)

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
		fmt.Println(arr[i])
	}

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	d2 := [2][2]int{{1, 2}, {3, 4}}

	for i := 0; i < len(d2); i++ {
		for j := 0; j < len(d2[0]); j++ {
			fmt.Println("NUM:=", d2[i][j])
		}
	}

	fmt.Println("===== Slices starts =====") // slice starts

	var slice []int // 1st way to create slice
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println("Slice:", slice)

	data := append(slice[1:2])
	fmt.Println("Data:", data)

	test := make([]int, 5)
	length := len(test)
	test = append(test[length:])
	fmt.Println(test)

	test = append(test, 6, 2, 1, 5, 4)

	test = append(test[:2], test[3:length]...)

	fmt.Println(sort.IntsAreSorted(test))

	sort.Ints(test)
	fmt.Println("TEST:", test)
}
