package Quest10

import("github.com/01-edu/z01")

func DescendComb() {
 for i := 99; i >= 01; i-- {
		for j := i-1; j >= 00; j-- {
			if i > j {
				z01.PrintRune(rune(i/10 + '0'))
				z01.PrintRune(rune(i%10 + '0'))
				z01.PrintRune(' ')
				z01.PrintRune(rune(j/10 + '0'))
				z01.PrintRune(rune(j%10 + '0'))
				if i != 01 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

			}
		}
	}
}