package helper

import (
	"fmt"
	"os"
	"sync"
)

func LengthOfString(wg *sync.WaitGroup, str string) int {
	defer wg.Done()
	fmt.Println("Get PID LengthOfString:", os.Getpid())
	return len(str)
}
