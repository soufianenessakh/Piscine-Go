package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	file, err := os.Open(args)
	if err != nil {
		fmt.Println("almost there!!")
		return
	}
	scanner := bufio.newscanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
