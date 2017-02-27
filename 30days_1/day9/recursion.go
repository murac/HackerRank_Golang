package main

import (
	"fmt"
)

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	var input int
	fmt.Scanf("%d", &input)
	res := factorial(input)
	fmt.Println(res)
}