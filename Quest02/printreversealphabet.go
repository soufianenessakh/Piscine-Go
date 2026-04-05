package piscine

import "github.com/01-edu/z01"

func Printreversealphabet() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
}
