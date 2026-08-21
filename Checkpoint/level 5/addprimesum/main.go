package main

import("github.com/01-edu/z01"
				"os"           )

func Atoi(s string)int{
	n:=0
	for _,c:=range(s){
		if c<'0'|| c>'9'{
			return -1
		}
		n=n*10+int(c-'0')
	}
	return n
}

func PrintNbr(n int) {
	if n==0{
		z01.PrintRune('0')
		return
	}
	if n<0{
	z01.PrintRune('-')
	n=-n
	}
	var digits [] int
	for n> 0{
		digits= append(digits,n%10)
		n/=10 
	}
	for i:= len (digits)-1;i>=0; i--{
		z01.PrintRune(rune(digits[i]+'0'))
	}
}
func IsPrime(nb int)bool{
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func main(){
		if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	n := Atoi(os.Args[1])

	if n <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := 0

	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}

	PrintNbr(sum)
	z01.PrintRune('\n')

}
	