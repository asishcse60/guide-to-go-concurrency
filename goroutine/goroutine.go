package goroutine

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	for i:=0; i<10; i++{
		value := rand.Intn(10)
		fmt.Printf("Value calculated: %v\n", value)
		values <- value
	}
}