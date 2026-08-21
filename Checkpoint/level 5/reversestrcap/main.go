package main

import("github.com/01-edu/z01"
				"os"			)

func main(){
	args:=os.Args[1:]
	for i:=0;i<len(args);i++{
		arg:=args[i]
		for j:=0;j<len(arg);j++{
			c:=arg[j]
			if c==' '{
				z01.PrintRune(' ')
				continue
			}else if len(arg)-1==j || arg[j+1]==' '  {
				if c>='a'&& c<='z'{
					c=c-32
				}
			}else{
				if c>='A' && c<='Z'{
					c=c+32
				}
			}
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
	}
}	