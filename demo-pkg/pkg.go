package main

import (
	"fmt"

	"example.com/pkg/hello"
)

func main()  {
	fmt.Println(hello.Hello("Perry"))
	fmt.Println(hello.Hi("Perry"))
}
