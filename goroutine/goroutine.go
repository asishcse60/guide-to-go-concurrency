package goroutine

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Printf("Value calculated: %v\n", value)
	values <- value
}