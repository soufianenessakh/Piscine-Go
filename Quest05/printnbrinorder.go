package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	count := []int{}
	for n > 0 {
		digit := n % 10
		count = append(count, digit)
		n = n / 10
	}
	for i := 0; i < len(count)-1; i++ {
		if count[i] > count[i+1] {
			temp := count[i]
			count[i] = count[i+1]
			count[i+1] = temp
		}
	}
	for _, nbr := range count {
		z01.PrintRune(rune(nbr + '0'))
	}
}
