package main

import "fmt"

func main() {
	fmt.Println("===== Struct Demo ======")

	user := User{"Savan", 21}
	fmt.Println("User data:", user)

	fmt.Printf("%+v\n", user)

	fmt.Println("===== Functions Demo =====")
	res1 := add(5, 6)
	fmt.Printf("Add: %v\n", res1)

	res2 := addAll(1, 2, 3, 4, 50, 6)
	fmt.Printf("AddAll: %v\n", res2)

	res3, msg := addWithMsg(1, 2, 3, 4, 5, 6)
	fmt.Printf("AddWithMsg: %v\n", res3)
	fmt.Printf("AddWithMsg Msg: %v\n", msg)
}

func add(a int, b int) int {
	return a + b
}

func addAll(values ...int) int {
	sum := 0
	for _, val := range values {
		fmt.Println(val)
		sum += val
	}
	return sum
}

func addWithMsg(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum, "This msg from AddWithMsg formate!!"
}

type User struct {
	Name string
	Age  int
}
