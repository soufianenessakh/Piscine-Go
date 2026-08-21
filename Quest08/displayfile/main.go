package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	if len(os.Args) < 2 {
    	return
	}
	args := os.Args[1]
	file, err := os.Open(args)
	if err != nil {
		fmt.Println("almost there!!")
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
    	fmt.Println("error reading file")
    	return
	}
	fmt.Print(string(content))
}
