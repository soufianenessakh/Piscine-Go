package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printRuneStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	printRuneStr("--insert")
	printRuneStr("  -i")
	printRuneStr("         This flag inserts the string into the string passed as argument.")
	printRuneStr("--order")
	printRuneStr("  -o")
	printRuneStr("         This flag will behave like a boolean, if it is called it will order the argument.")
}

func bubbleSortRunes(runes []rune) {
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
}
func main() {
	args := os.Args[1:]

	insert := false
	order := false
	help := false
	insertValue := ""
	mainString := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insert = true
			insertValue = ""
			for _, r := range arg[9:] {
				insertValue += string(r)
			}
			continue
		}

		if arg == "-i" || arg == "--insert" {
			insert = true
			if i+1 < len(args) {
				insertValue = ""
				for _, r := range args[i+1] {
					insertValue += string(r)
				}
				i++
			} else {
				printRuneStr("Error: --insert flag needs a value")
				return
			}
		} else if arg == "-o" || arg == "--order" {
			order = true
		} else if arg == "-h" || arg == "--help" {
			help = true
		} else {
			mainString = ""
			for _, r := range arg {
				mainString += string(r)
			}
		}
	}
	if help || len(args) == 0 {
		printHelp()
		return
	}
	if insert && insertValue != "" {
		mainString += insertValue
	}
	if order && mainString != "" {
		runes := []rune{}
		for _, r := range mainString {
			runes = append(runes, r)
		}
		bubbleSortRunes(runes)
		mainString = ""
		for _, r := range runes {
			mainString += string(r)
		}
	}

	printRuneStr(mainString)
}