package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inputFile = "day04/input-0"
	X         = 88
	M         = 77
	A         = 65
	S         = 83
	up        = iota
	down
	left
	right
	rightUp
	leftDown
	leftUp
	rightDown
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	grid := make([][]int32, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		gridLine := make([]int32, len(line))
		for i, c := range line {
			gridLine[i] = c
		}
		grid = append(grid, gridLine)
	}

	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for dir := up; dir <= rightDown; dir++ {
				if search(grid, i, j, dir, []int32{}) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	grid := make([][]int32, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		gridLine := make([]int32, len(line))
		for i, c := range line {
			gridLine[i] = c
		}
		grid = append(grid, gridLine)
	}

	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == A && isCrossMAS(grid, i, j) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func numToChar(num int32) string {
	switch num {
	case X:
		return "X"
	case M:
		return "M"
	case A:
		return "A"
	case S:
		return "S"
	}
	return ""
}

func isCrossMAS(grid [][]int32, i, j int) bool {
	if i-1 < 0 || i+1 >= len(grid) || j-1 < 0 || j+1 >= len(grid[i]) {
		return false
	}

	// M.S
	// .A.
	// M.S
	if grid[i-1][j-1] == M && grid[i+1][j-1] == M && grid[i-1][j+1] == S && grid[i+1][j+1] == S {
		return true
	}
	// S.S
	// .A.
	// M.M
	if grid[i-1][j-1] == S && grid[i+1][j-1] == M && grid[i-1][j+1] == S && grid[i+1][j+1] == M {
		return true
	}
	// M.M
	// .A.
	// S.S
	if grid[i-1][j-1] == M && grid[i+1][j-1] == S && grid[i-1][j+1] == M && grid[i+1][j+1] == S {
		return true
	}
	// S.M
	// .A.
	// S.M
	if grid[i-1][j-1] == S && grid[i+1][j-1] == S && grid[i-1][j+1] == M && grid[i+1][j+1] == M {
		return true
	}

	return false
}

func isXMAS(path []int32) bool {
	if len(path) != 4 {
		return false
	}

	for i, c := range []int32{X, M, A, S} {
		if path[len(path)-i-1] != c {
			return false
		}
	}

	return true
}

func search(grid [][]int32, i, j, dir int, path []int32) bool {
	if isXMAS(path) {
		return true
	}

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || len(path) > 4 {
		return false
	}

	switch dir {
	case up:
		return search(grid, i-1, j, up, append(path, grid[i][j]))
	case down:
		return search(grid, i+1, j, down, append(path, grid[i][j]))
	case left:
		return search(grid, i, j-1, left, append(path, grid[i][j]))
	case right:
		return search(grid, i, j+1, right, append(path, grid[i][j]))
	case rightUp:
		return search(grid, i-1, j+1, rightUp, append(path, grid[i][j]))
	case leftDown:
		return search(grid, i+1, j-1, leftDown, append(path, grid[i][j]))
	case leftUp:
		return search(grid, i-1, j-1, leftUp, append(path, grid[i][j]))
	case rightDown:
		return search(grid, i+1, j+1, rightDown, append(path, grid[i][j]))
	}

	return false
}
