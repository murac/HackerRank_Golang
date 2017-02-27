package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var input, number int
	var name string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d", &input)
	phone_book := make(map[string]int, input)
	for x := 0; x < input; x++ {
		fmt.Scanf("%s %d", &name, &number)
		phone_book[name] = number
	}
	for scanner.Scan() {
		name = scanner.Text()
		res := phone_book[name]
		if res != 0 {
			fmt.Printf("%s=%d\n", name, res)
		} else {
			fmt.Println("Not found")
		}
	}
}
