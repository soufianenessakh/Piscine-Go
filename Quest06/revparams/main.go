package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0 ; i-- {
		for _, arg := range args[i] {
			z01.PrintRune(arg)
		}
		z01.PrintRune('\n')
	}
}
