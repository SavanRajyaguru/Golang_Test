package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("====== File Start ======")

	content := "This is Go lang test code :)"

	file, err := os.Create("./MyText.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("File Data Length:", length)
	defer file.Close()

	readFile("./MyText.txt")

}

func readFile(filepath string) {
	dataByte, err := os.ReadFile(filepath)
	checkNilErr(err)
	fmt.Println("File Data:", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
