package main

import "fmt"

func main() {
	// =============== variables ============= //
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// =============== Loops ============= //
	// for is Go’s only looping construct. Here are some basic types of for loops.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// infinite loop
	for {
		fmt.Println("loop")
		break
	}

	// =============== conditions ============= //
	if num := 11; num > 0 {
		fmt.Println(num)
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
