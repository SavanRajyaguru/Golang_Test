package helper

import (
	"fmt"
	"os"
	"sync"
)

func SumOfVal(wg *sync.WaitGroup, values ...int) int {
	defer wg.Done()
	fmt.Println("Get PID SumOfVal:", os.Getpid())
	sum := 0
	for _, val := range values {
		sum += val
	}
	fmt.Println("Sum:", sum)
	return sum
}
