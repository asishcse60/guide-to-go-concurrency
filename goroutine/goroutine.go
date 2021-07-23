package goroutine

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
)

func CalculateValue(values chan int) {
	for i:=0; i<10; i++{
		value := rand.Intn(10)
		fmt.Printf("Value calculated: %v\n", value)
		values <- value
	}
}

func Fetch( url string, wg *sync.WaitGroup){
	res, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("Url download: %v and Code: %v\n", url, res.StatusCode)
	wg.Done()
}