package main

import("github.com/01-edu/z01"
				"os"				)

func main(){
	if len(os.Args)!=3{
		z01.PrintRune('\n')
		return
	}
	arg:=os.Args[1]
	arg1:=os.Args[2]
	c:=make(map[rune]bool)
	for _,ch:=range(arg) {
		if !c[ch]{
			c[ch]=true
			z01.PrintRune(ch)
		}
	}
	for _,cha:=range(arg1){
			if !c[cha]{
				c[cha]=true
				z01.PrintRune(cha)
			}
	}
	z01.PrintRune('\n')
}