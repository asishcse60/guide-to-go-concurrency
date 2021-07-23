package main

import (
	"fmt"
	"runtime"
)

func main() {
	//variable declaration
	//var welcome string="Hello World!"
	welcome := "Hello World!"
	fmt.Println(welcome)

	var myAge = 27
	fmt.Printf("Age is %v\n", myAge)
	//Constant declaration: const name type(optional) = value
	const pi  = 3.1416
	fmt.Printf("Pi Value is: %v\n", pi)
	//println("Hello")
	CheckCustomerHeight(41)
	CheckCustomerHeight(18)
	CheckCustomerHeight(10)
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