package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value: ")

	// comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Input: ", input)

	conv, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Conv: %v And Type: %T", conv, conv)
	}
}
