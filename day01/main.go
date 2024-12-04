package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("day01/input-0")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	lq := make([]int, 0)
	rq := make([]int, 0)

	s := bufio.NewScanner(f)
	var l, r int
	for s.Scan() {
		_, _ = fmt.Sscanf(s.Text(), "%d %d", &l, &r)
		lq = append(lq, l)
		rq = append(rq, r)
	}

	part1(lq, rq)
	part2(lq, rq)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func part1(lq, rq []int) {
	sort.Ints(lq)
	sort.Ints(rq)
	var total int
	for i := 0; i < len(lq); i++ {
		total += abs(rq[i] - lq[i])
	}
	fmt.Println(total)
}

func part2(lq, rq []int) {
	var total int
	for i := 0; i < len(lq); i++ {
		var count int
		for j := 0; j < len(lq); j++ {
			if lq[i] == rq[j] {
				count++
			}
		}
		total += lq[i] * count
	}
	fmt.Println(total)
}
