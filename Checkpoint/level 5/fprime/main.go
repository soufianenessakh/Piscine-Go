package main

import (
	"github.com/01-edu/z01"
	"os"
)

func PrintNbr(n int) {
	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	num := 0

	for _, c := range args {
		if c < '0' || c > '9' {
			return
		}
		var digit int
		digit = int(c - '0')
		num = num*10 + digit
	}

	if num <= 1 {
		return
	}

	first := true

	for i := 2; num > 1; {
		if num%i == 0 {
			if !first {
				z01.PrintRune('*')
			}
			PrintNbr(i)
			num = num / i
			first = false
		} else {
			i++
		}
	}
	z01.PrintRune('\n')
}