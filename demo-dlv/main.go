package main

import (
	"fmt"
	"time"
)

var k string

func main() {
	k = "v"
	fmt.Println("Hello, World!")
	time.Sleep(time.Second * 3)
	fmt.Println("Hello, Again")
}
