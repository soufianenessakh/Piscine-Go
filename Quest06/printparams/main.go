package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	agrs := os.Args[1:]
	for _, arg := range agrs {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
