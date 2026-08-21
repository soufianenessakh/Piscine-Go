package main

import("github.com/01-edu/z01"
			"os"
)
func main(){
	if len(os.Args)!=2{
		return
	}
	hasword:=false
	inword:=false
	needspace:=false
	s:=os.Args[1]
	for _,c:=range s{
		if c!=' ' && c!='\t' {
			if needspace==true{
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				needspace=false
			}
			inword=true
			hasword=true
			z01.PrintRune(c)
		} else {
			if inword == true {
				needspace = true
				inword = false
			}
		}
	}
	if hasword==true{
		z01.PrintRune('\n')
	}
}