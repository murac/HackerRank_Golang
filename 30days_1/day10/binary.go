package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int64
	var count,max int
	fmt.Scanf("%d", &input)
	bin := strconv.FormatInt(input,2)
	for x:=0;x<len(bin);x++{
		if bin[x] == '1' {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(max)
}