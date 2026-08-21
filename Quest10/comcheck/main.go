package main

import("github.com/01-edu/z01"
				"os"
)

func main(){
	args:=os.Args[1:]
	for i:=0;i<len(args);i++{
		if args[i]=="01" || args[i]=="galaxy" || args[i]=="galaxy 01"{
			z01.PrintRune('A')
			z01.PrintRune('l')
			z01.PrintRune('e')
			z01.PrintRune('r')
			z01.PrintRune('t')
			z01.PrintRune('!')
			z01.PrintRune('!')
			z01.PrintRune('!')
			z01.PrintRune('\n')
			return
		}
	}
}