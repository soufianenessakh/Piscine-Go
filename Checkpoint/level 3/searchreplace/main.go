package main

import("github.com/01-edu/z01"
			"os"
)

func main(){
	if len(os.Args)!=4{
		return
	}
	str:=os.Args[1]
	old:=os.Args[2]
	new:=os.Args[3]
	if len(old)!=1||len(new)!=1{
		return
	}
	for i:=0;i<len(str);i++{
		if str[i]==old[0]{
			z01.PrintRune(rune(new[0]))
		}else{
			z01.PrintRune(rune(str[i]))
		}
	}
	z01.PrintRune('\n')
}