package main

import (
	"fmt"
)

func main() {
	var input int
	var cur_str, cur_let, odd, even string
	fmt.Scanf("%d", &input)
	words := make([]string, input)
	for x := 0; x < input; x++ {
		fmt.Scanln(&cur_str)
		words[x] = cur_str
		for i := 0; i < len(cur_str); i++ {
			cur_let = string(cur_str[i])
			if i%2 == 0 {
				even += cur_let
			} else {
				odd += cur_let
			}
		}
		words[x] = even + " " + odd
		even = ""
		odd = ""
	}
	for x := 0; x < input; x++ {
		fmt.Println(words[x])
	}
}
