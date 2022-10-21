package main

import "fmt"

func main() {
	fmt.Printf("%s", string([]byte{}))
	fmt.Printf("%s", string([]byte("")))
}
