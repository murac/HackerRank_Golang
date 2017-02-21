package main

import "fmt"

func main() {
	var input,prod int
	fmt.Scanf("%d",&input)
	for x:=1;x<=10;x++ {
		prod = input * x
		fmt.Printf("%d x %d = %d\n",input,x,prod)
	}
}
