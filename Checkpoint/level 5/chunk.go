package checkpoint

import("github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}


func Chunk(slice []int, size int) {
	if size==0 {
		z01.PrintRune('\n')
		return 
	}
	var output [][]int
	for i:=0;i<len(slice) ;i+=size{
		end:=i+size
		if end>len(slice){
			end=len(slice)
		}
		chunk:=slice[i:end]
		output=append(output, chunk)
	}
	z01.PrintRune('[')

	for i, chunk := range output {
		z01.PrintRune('[')

		for j, num := range chunk {
			PrintNbr(num)
			if j != len(chunk)-1 {
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune(']')

		if i != len(output)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune(']')
	z01.PrintRune('\n')
}