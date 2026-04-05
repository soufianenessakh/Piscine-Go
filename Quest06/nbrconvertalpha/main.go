package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	index := 0
	for _, arg := range args {
		if arg[0] == "--upper" {
			upper = true
			index = 1
		}
		for i := index; i <= len(arg); i++ {
			err, num := strconv.Atoi(arg[i])
			if err !=nil || num < 1 || num > 26 {
				z01.PrintRune(' ')
			}else{
				if upper
				z01.PrintRune(rune(num + 64))
			}else{
			z01.PrintRune(rune(num + 96))
			}
		}
	}
	z01.PrintRune('\n')
}