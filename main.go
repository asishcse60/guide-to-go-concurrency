package main

import (
	"fmt"
	"github.com/asishcse60/guide-to-go-concurrency/goroutine"
	"runtime"
	"time"
)
type Employee interface {
	GetName() string
}
type Engineer struct {
	Name string
}

type Manager struct {
	Name string
}

func (e *Engineer)	GetName()string  {
	return "engineer name: " + e.Name
}

func (m *Manager)	GetName()string  {
	return "Manager name: " + m.Name
}
func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}


func main() {

	//Channel
	//values := make(chan int)
	//go goroutine.CalculateValue(values)
	//value := <-values
	//fmt.Println(value)
	//Buffers channel
	values := make(chan int, 3)
	go goroutine.CalculateValue(values)
	for i := 0; i < 10; i++ {
		time.Sleep(1 *time.Second)
		value:= <-values
		fmt.Println(value)
	}

	//engineer := &Engineer{Name: "Ashish"}
	//manager := &Manager{Name: "Boss"}

	//PrintDetails(engineer)
	//PrintDetails(manager)
	//variable declaration
	//var welcome string="Hello World!"
	//HelloWorld("Ashish", 27, 6)
	//welcome := "Hello World!"
	//fmt.Println(welcome)

	//var myAge = 27
	//fmt.Printf("Age is %v\n", myAge)
	//Constant declaration: const name type(optional) = value
	//const pi  = 3.1416
	//fmt.Printf("Pi Value is: %v\n", pi)
	//println("Hello")
	//CheckCustomerHeight(41)
	//CheckCustomerHeight(18)
	//CheckCustomerHeight(10)

	//ages := []int{20,25,30,35,40}
	//for i := 0; i < len(ages); i++ {
		//fmt.Printf("age %v is %v\n", i, ages[i])
	//}
	///fmt.Println("Other for loop run")
	//for i,k := range ages	{
		//fmt.Printf("age %v is %v\n", i, k)
	//}
}

func HelloWorld(name string, age, height int) {
	fmt.Println("Hello", name)
	fmt.Println("Age", age)
	fmt.Println("Height", height)
}
func CheckCustomerHeight(height int) {

	switch os := runtime.GOOS; os {
	case "darwin":
		println("os X")
	case "linux":
		println("linux machine")
	case "windows":
		println("windows machine")
	default:
		println("something else")
	}

	switch  {
	case height>=18 && height<=40:
		println("Young")
	case height>=41:
		println("Old")
	default:
		println("Under Age")
	}

	switch  height{
	case 18:
		println("Young: 18")
		break
	case 41:
		println("Old: 41")
		break
	default:
		println("Under Age: <18")
	}
}