package main

import("github.com/01-edu/z01"
				"os"			)

func main(){
	if len(os.Args)!=2{
		z01.PrintRune('\n')
		return
	}

	args:=os.Args[1]
	end:=len(args)-1
	for i:=len(args)-1;i>=0;i--{
		if args[i]==' '{
			for j:=i+1;j<=end;j++{
				z01.PrintRune(rune(args[j]))
			}
			z01.PrintRune(' ')
			end=i-1
		}
	}
	for j:=0;j<=end;j++{
		z01.PrintRune(rune(args[j]))
	}
	z01.PrintRune('\n')
}