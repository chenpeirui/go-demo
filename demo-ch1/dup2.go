package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	if len(os.Args) > 1 {
		files := os.Args[1:]
		for file := range files {
			file, err := os.Open(file)
			if err != nil {
				fmt.Errorf(err)
			}

			input := bufio.NewScanner(file)
			countLines(lines, counts)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(
