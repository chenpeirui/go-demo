package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	sayHello()
	a := "a"
	time.Sleep(time.Second)
	fmt.Println("Hello, Again!")
	fmt.Println(a)
}
