package main

import (
	"fmt"
	"math/rand"
)

const (
	row = 5
	col = 4
)

func rowLine() {
	fmt.Println()
	for i := 0; i < col; i++ {
		fmt.Print(" -----")
	}
	fmt.Println()
}

func countLiveNeighbourCell(a [row][col]int, r, c int) int {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i == r && j == c) || i < 0 || j < 0 || i >= row || j >= col {
				continue
			}
			if a[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	var a [row][col]int
	var b [row][col]int

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			a[i][j] = rand.Intn(2)
		}
	}

	fmt.Println("Initial Stage:")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ", a[i][j])
		}
		fmt.Printf("\n")
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			neighbourLiveCell := countLiveNeighbourCell(a, i, j)
			if a[i][j] == 1 && (neighbourLiveCell == 2 || neighbourLiveCell == 3) {
				b[i][j] = 1
			} else if a[i][j] == 0 && neighbourLiveCell == 3 {
				b[i][j] = 1
			} else {
				b[i][j] = 0
			}
		}
	}

	fmt.Println("\nNext Generation:")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ", b[i][j])
		}
		fmt.Printf("\n")
	}
}
