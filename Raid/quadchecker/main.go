package main

import (
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	fmt.Println(input)
}