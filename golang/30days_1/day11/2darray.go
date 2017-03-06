package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	grid := [6][6]int{}
	scanner := bufio.NewScanner(os.Stdin)
	max := -100
	for idx := 0; scanner.Scan() && idx < 6; idx++ {
		cur := strings.Split(scanner.Text(), " ")
		for x := 0; x < len(cur); x++ {
			number, _ := strconv.Atoi(cur[x])
			grid[idx][x] = number
		}
	}
	//fmt.Println(grid)
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			cur := grid[x][y] +
				grid[x][y+1] +
				grid[x][y+2] +
				grid[x+1][y+1] +
				grid[x+2][y] +
				grid[x+2][y+1] +
				grid[x+2][y+2]
			if cur > max {
				max = cur
			}
		}
	}
	fmt.Println(max)
}