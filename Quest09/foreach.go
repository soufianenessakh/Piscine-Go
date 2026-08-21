package Quest09

import("github.com/01-edu/z01")

func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
}
func PrintNbr(n int) {
    if n == 0 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }

    var digits []rune
    for n > 0 {
        digits = append([]rune{rune(n%10 + '0')}, digits...)
        n /= 10
    }

    for _, d := range digits {
        z01.PrintRune(d)
    }
    z01.PrintRune('\n')
}
