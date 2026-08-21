package main

import("github.com/01-edu/z01"
				"os"           )

func main(){
	if len(os.Args)!=3{
		return
	}
	s1:=os.Args[1]
	s2:=os.Args[2]
	i:=0
		for _,c:=range(s2){
			if i>=len(s1){
				break
			}
			if c==rune(s1[i]){
				i++
			}
		}
		if i==len(s1) {
			z01.PrintRune('1')
		}else{
			z01.PrintRune('0')
		}
}