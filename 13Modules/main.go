package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}

type User struct {
	Name string
	Age  int
}

// Anonymous Struct Fields
type student struct {
	int
	string
	float64
	user User
}

func (u *User) setData() string {
	return u.Name + " " + strconv.Itoa(u.Age)
}

func main() {
	fmt.Println("====== My Modules ======")
	// user := User{"Savan", 21}

	slice := make([]int, 2, 4)
	// arr1 := make([]int, 5)
	arr := []int{1, 2, 3, 4, 5}
	slice = append(slice, arr...)

	fmt.Println("Main function started.")
	doSomething()
	fmt.Println("Main function continues.")

	fmt.Println("Data:", slice)
	a := []User{{"asdfs", 21}, {"werwer", 22}, {"qrwer", 21}}
	// b := []User{{"aaa", 21}, {"bbb", 22}, {"ccc", 21}}
	a = append(a, User{"saaaaa", 22})

	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Age > a[j].Age
	})

	fmt.Println("A>>>>", a)
	fmt.Println("SUM>>>", test("asdf", a...))
	fmt.Println("a>>>>", a)
	fmt.Println("Time: ", time.Now().Unix())
	tm, _ := strconv.ParseInt("1695722067", 10, 64)
	fmt.Println("CURR: ", time.Unix(tm, 0))

	value1 := data(20)
	value2 := data(20)
	value3 := data(20)
	res := value1.multiply(value2.multiply(value3))
	fmt.Println("Final result: ", res)

	fmt.Printf("%+v", student{21, "asdf", 1.2, User{"lalo", 22}})
}

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func test(c string, u ...User) int {
	sum := 0
	for _, data := range u {
		fmt.Println("VAL:", data.Age)
	}
	return sum
}

func doSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Doing something...")
	panic("Oops, an error occurred!")
}
