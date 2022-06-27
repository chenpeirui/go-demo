package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func v1() string {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	return s
}

func v2() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	// repeat count
	n := 10000000

	// run v1()
	start := time.Now()
	for i := 0; i < n; i++ {
		v1()
	}
	cost1 := time.Since(start)

	// run v2()
	start = time.Now()
	for i := 0; i < n; i++ {
		v2()
	}
	cost2 := time.Since(start)

	// compare cost
	fmt.Println(cost1 - cost2)
}
