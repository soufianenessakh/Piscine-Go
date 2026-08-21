package Quest10

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for i, card := range deck {

		if i%3 == 0 {
			z01.PrintRune('P')
			z01.PrintRune('l')
			z01.PrintRune('a')
			z01.PrintRune('y')
			z01.PrintRune('e')
			z01.PrintRune('r')
			z01.PrintRune(' ')
			z01.PrintRune(rune(i/3 + 1 + '0'))
			z01.PrintRune(':')
		}
		if card >= 10 {
			z01.PrintRune(rune(card/10 + '0'))
		}
		z01.PrintRune(rune(card%10 + '0'))
		z01.PrintRune(' ')
		if i%3 == 2 {
			z01.PrintRune('\n')
		}
	}
}