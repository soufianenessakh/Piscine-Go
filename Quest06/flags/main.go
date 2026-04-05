package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	flag := false
	nbrflags := 0

	for _, arg := range args {
		if arg == "--order" || arg == "-o" {
			flag = true
			nbrflags
		}
		if arg == "--insret" || arg == "-i" {
			flag = true
			nbrflags
		}
		if arg == "--help" || arg == "-h" {
			flag = true
			nbrflags
		}

	}
}
