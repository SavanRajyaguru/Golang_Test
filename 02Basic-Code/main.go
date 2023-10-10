package main

import (
	"fmt"
	"math/rand"
)

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

	// goto statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			// statement is always upper line from the declare if not then is go infinite
			goto loc
		}
		fmt.Println("Number>>>>", i)
	}
loc:
	fmt.Println("This is the goto statement")

	// infinite loop
	for {
		fmt.Println("loop")
		break
	}

	// =============== conditions ============= //
	if num := 11; num > 0 {
		fmt.Println(num)
	}

	num := rand.Intn(5) + 1
	fmt.Println("Random:", num)

	switch num {
	case 1:
		fmt.Println("1 is there")
	case 2:
		fmt.Println("2 is there")
	case 3:
		fmt.Println("3 is there")
	default:
		fmt.Println("Data is not")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool", t)
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
