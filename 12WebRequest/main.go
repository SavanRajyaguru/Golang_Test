package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "httpss://loc.dev"

func main() {
	fmt.Println("====== Web Request Start ======")

	res, err := http.Get(url)

	checkNilErr(err)
	fmt.Printf("Response Type: %T\n", res)

	defer res.Body.Close()

	dataByte, err := io.ReadAll(res.Body)

	checkNilErr(err)

	fmt.Println("Data>>>>>>>>>>>>>>>:", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
