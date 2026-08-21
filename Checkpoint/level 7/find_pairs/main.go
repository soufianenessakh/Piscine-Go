package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arr := os.Args[1]
	target := (os.Args[2])
	if len(target)>1{
		fmt.Println("Invalid target sum.")
		return
	}

	if len(arr) < 2 || arr[0] != '[' || arr[len(arr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	numbers := []int{}

	for i := 1; i < len(arr)-1; i++ {
		num := string(arr[i])
		if num == " " || num == "," {
			continue
		}
		if '0' <= arr[i] && arr[i] <= '9' {
			numbers = append(numbers, Atoi(num))
		} else {
			fmt.Println("Invalid input.")
			return
		}
		if Atoi(num) < 0 {

		}
	}

	output := [][]int{}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == Atoi(target) {
				output = append(output, []int{i, j})
			}
		}
	}

	if len(output) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, output)
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sign := 1
	output := 0
	star := 0
	if s[0] == '-' {
		sign = -1
		star = 1
	} else if s[0] == '+' {
		star = 1
	}
	if star == len(s) {
		return 0
	}
	for i := star; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		output = output*10 + (int(s[i]) - '0')
	}
	return output * sign
}
