package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 2 {
        z01.PrintRune('\n')
        return
    }
    args := os.Args[1]
    output:=""
    for i := 0; i < len(args); i++ {
        if args[i] != ' ' && args[i + 1 ] == ' '  && output == "" {
            for j:=i+1;j<len(args);j++{

                if args[j]==' ' && args[j+1]==' '{
                    continue
                }
                 output+=string(args[j])
            }
            break
        }else {
            continue
        }
    }
    
    output += " "

    start := false
    for _, c := range args {
    if c == ' ' && !start {
        continue
    }

    if c == ' ' && start {
        break
    }

    start = true
    output += string(c)
}
    for _,c:=range(output){
        z01.PrintRune(c)
    }
    z01.PrintRune('\n')
}