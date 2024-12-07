package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	dot      = 46
	obstacle = 35
	pos      = 94
	up       = 0 // ^
	right    = 1 // >
	down     = 2 // v
	left     = 3 // <
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day06/input-0")
	s := bufio.NewScanner(f)

	x, y := 0, 0
	grid := make([][]int32, 0)
	for s.Scan() {
		row := make([]int32, 0)
		for _, c := range s.Text() {
			if c == pos {
				x, y = len(grid), len(row)
			}
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	walk(grid, x, y, up)

	count := 0
	for _, row := range grid {
		for _, c := range row {
			if c == 99 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	f, _ := os.Open("day06/input-0")
	s := bufio.NewScanner(f)

	x, y := 0, 0
	grid := make([][]int32, 0)
	for s.Scan() {
		row := make([]int32, 0)
		for _, c := range s.Text() {
			if c == pos {
				x, y = len(grid), len(row)
			}
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	path := make([][3]int, 0)
	findLoop(grid, x, y, up, &path)

	uniquePathMap := make(map[[2]int]struct{})
	for _, p := range path {
		uniquePathMap[[2]int{p[0], p[1]}] = struct{}{}
	}
	uniquePath := make([][2]int, 0)
	for k := range uniquePathMap {
		uniquePath = append(uniquePath, k)
	}

	var count int
	for _, p := range uniquePath {
		grid[p[0]][p[1]] = obstacle
		if findLoop(grid, x, y, up, &[][3]int{}) {
			count++
		}
		grid[p[0]][p[1]] = dot
	}
	fmt.Println(count)
}

func detectLoop(path [][3]int) bool {
	last := path[len(path)-1]
	for i := len(path) - 2; i >= 0; i-- {
		if path[i] == last {
			return true
		}
	}
	return false
}

func findLoop(grid [][]int32, x, y, dir int, path *[][3]int) bool {
	*path = append(*path, [3]int{x, y, dir})
	if isOutOfBounds(grid, x, y, dir) {
		return false
	}

	for isObstacle(grid, x, y, dir) {
		dir = (dir + 1) % 4
	}

	if detectLoop(*path) {
		return true
	}

	switch dir {
	case up:
		return findLoop(grid, x-1, y, dir, path)
	case right:
		return findLoop(grid, x, y+1, dir, path)
	case down:
		return findLoop(grid, x+1, y, dir, path)
	case left:
		return findLoop(grid, x, y-1, dir, path)
	}

	return false
}

func printLoop(grid [][]int32, path [][3]int) {
	last := path[len(path)-1]
	grid[last[0]][last[1]] = 60
	for i := len(path) - 2; i >= 0; i-- {
		if last == path[i] {
			break
		}
		grid[path[i][0]][path[i][1]] = 61
	}
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	for i := len(path) - 1; i >= 0; i-- {
		grid[path[i][0]][path[i][1]] = dot
	}
	fmt.Println()
	fmt.Println()
}

func walk(grid [][]int32, x, y, dir int) {
	grid[x][y] = 99
	if isOutOfBounds(grid, x, y, dir) {
		return
	}

	if isObstacle(grid, x, y, dir) {
		dir = (dir + 1) % 4
	}

	switch dir {
	case up:
		walk(grid, x-1, y, dir)
	case right:
		walk(grid, x, y+1, dir)
	case down:
		walk(grid, x+1, y, dir)
	case left:
		walk(grid, x, y-1, dir)
	}
}

func isOutOfBounds(grid [][]int32, x, y, dir int) bool {
	switch dir {
	case up:
		return x-1 < 0
	case right:
		return y+1 >= len(grid[0])
	case down:
		return x+1 >= len(grid)
	case left:
		return y-1 < 0
	}
	return false
}

func isObstacle(grid [][]int32, x, y, dir int) bool {
	switch dir {
	case up:
		return grid[x-1][y] == obstacle
	case right:
		return grid[x][y+1] == obstacle
	case down:
		return grid[x+1][y] == obstacle
	case left:
		return grid[x][y-1] == obstacle
	}
	return false
}
