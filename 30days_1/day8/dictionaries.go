package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	nums := make([]int, input)
	for x := 0; x < input; x++ {
		fmt.Scanf("%d", &nums[x])
	}
	for x := input - 1; x >= 0; x-- {
		fmt.Printf("%d ", nums[x])
	}
}