package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	i := 0
	j := 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

		for _, c := range s1 {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	
}