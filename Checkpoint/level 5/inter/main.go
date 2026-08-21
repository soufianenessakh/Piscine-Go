package main

import("github.com/01-edu/z01"
				"os"			)

func main(){
	if len(os.Args)!=3{
		return
	}
	s1:=os.Args[1]
	s2:=os.Args[2]
	firstrun:=make(map[rune]bool)
	for _,ch:=range(s1) {
		for _,c:=range(s2){
			if c==ch{
				if !firstrun[ch]{
					firstrun[ch]=true
					z01.PrintRune(ch)
				}
				break
			}

		}
	}
z01.PrintRune('\n')
}