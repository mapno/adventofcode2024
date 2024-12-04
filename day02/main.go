package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFile = "day02/input-0"

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var count int
	s := bufio.NewScanner(f)

	for s.Scan() {
		report := txtToReport(s.Text())
		if checkSafe(report) {
			count++
		}
	}
	fmt.Println(count)
}

func txtToReport(txt string) report {
	reportStr := strings.Split(txt, " ")
	report := make(report, len(reportStr))
	for i, r := range reportStr {
		report[i], _ = strconv.Atoi(r)
	}
	return report
}

func removeIdx(s []int, i int) []int {
	n := make([]int, len(s)-1)
	copy(n, s[:i])
	copy(n[i:], s[i+1:])
	return n
}

func part2() {
	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)

	var count int

	for s.Scan() {
		report := txtToReport(s.Text())
		var isSafe bool
		if checkSafe(report) {
			isSafe = true
		} else {
			for i := 0; i < len(report); i++ {
				if checkSafe(removeIdx(report, i)) {
					isSafe = true
				}
			}
		}
		if isSafe {
			count++
		}
	}
	fmt.Println(count)
}

func checkSafe(report report) bool {
	if !sort.IsSorted(report) && !sort.IsSorted(sort.Reverse(report)) {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		if diff := abs(report[i+1] - report[i]); diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

var _ sort.Interface = (*report)(nil)

type report []int

func (r report) Len() int { return len(r) }

func (r report) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r report) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
