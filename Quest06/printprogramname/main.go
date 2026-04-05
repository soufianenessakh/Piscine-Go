package main

import (
	"os"
	"github.com/01-edu/z01"
	"path/filepath"
)

func main() {
	args :=filepath.Base (os.Args[0]) 
	for _, char := range args {
			z01.PrintRune(char) 
	} 
}