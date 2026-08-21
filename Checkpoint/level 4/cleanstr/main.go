package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	inword := false
	needspace := false
	str := os.Args[1]

	for _, c := range str {
		if c != ' ' && c != '\t' {
			if needspace == true {
				z01.PrintRune(' ')
				needspace = false
			}
			z01.PrintRune(c)
			inword = true
		} else {
			if inword == true {
				needspace = true
				inword = false
			}
		}
	}

	z01.PrintRune('\n')
}