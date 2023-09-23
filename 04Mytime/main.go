package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	fmt.Println(time1)

	// format date
	fmt.Println(time1.Format("01/02/2006 Monday 15:04:05 MST"))

	// create date
	createDate := time.Date(2001, 12, 20, 12, 30, 12, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01/02/2006 Monday"))
}
