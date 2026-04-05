package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := 0; j <= 99; j++ {
			if i < j {
				z01.PrintRune(rune(i/10 + '0'))
				z01.PrintRune(rune(i%10 + '0'))
				z01.PrintRune(' ')
				z01.PrintRune(rune(j/10 + '0'))
				z01.PrintRune(rune(j%10 + '0'))
				if i != 98 {
					z01.PrintRune(';')
					z01.PrintRune(' ')
				}

			}
		}
	}
}
