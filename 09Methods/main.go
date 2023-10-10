package main

import "fmt"

func main() {
	fmt.Println("===== Methods Start =====")

	user := User{"savan", 21}
	user2 := User{"savan", 22}
	user.getAge()
	user2.UpdateAge()
	user2.getAge()
}

type User struct {
	Name string
	Age  int
}

func (u User) getAge() {
	fmt.Println("Get Age:", u.Age)
}

func (u User) UpdateAge() {
	u.Age = 25
	fmt.Println("Update Age:", u.Age)
}
