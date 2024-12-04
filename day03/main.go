package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const inputFile = "day03/input-0"

func main() {
	// part1()
	part2()
}

func part1() {
	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var total int
	for s.Scan() {
		for _, m := range r.FindAllString(s.Text(), -1) {
			var a, b int
			fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
			total += a * b
		}
	}
	fmt.Println(total)
}

func part2() {
	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)

	r := regexp.MustCompile(`(mul\(\d+,\d+\)|(do\(\)|don't\(\)))`)

	var total int
	do := true
	for s.Scan() {
		for _, m := range r.FindAllString(s.Text(), -1) {
			if m == "do()" {
				do = true
			} else if m == "don't()" {
				do = false
			} else {
				var a, b int
				fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
				if do {
					total += a * b
				}
			}
		}
	}
	fmt.Println(total)
}
