package checkpoint

import "github.com/01-edu/z01"

func printhexbyte(b byte) {
	hex := "0123456789abcdef"
	z01.PrintRune(rune(hex[b/16]))
	z01.PrintRune(rune(hex[b%16]))
}

func PrintMemory(arr [10]byte) {
	for i := 0; i < 10; i++ {
		printhexbyte(arr[i])

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != 9 {
			z01.PrintRune(' ')
		}
	}

	if 10%4 != 0 {
		z01.PrintRune('\n')
	}

	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}