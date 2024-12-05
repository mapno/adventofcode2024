package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day05/input-0"

func main() {
	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)

	rules := make(map[int][]int)
	updates := make([][]int, 0)
	for s.Scan() {
		var a, b int
		txt := s.Text()
		if _, err := fmt.Sscanf(txt, "%d|%d", &a, &b); err == nil {
			rules[a] = append(rules[a], b)
		} else if txt == "" {
			continue
		} else {
			update := make([]int, 0)
			for _, numStr := range strings.Split(txt, ",") {
				num, _ := strconv.Atoi(numStr)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	var validUpdates []int
	for i, update := range updates {
		valid := true
		for j, num := range update {
			rule := rules[num]
			for k := 0; k < j; k++ {
				if contains(rule, update[k]) {
					valid = false
					break
				}
			}
		}
		if valid {
			validUpdates = append(validUpdates, i)
		}
	}
	var count int
	for _, i := range validUpdates {
		count += middle(updates[i])
	}
	fmt.Println("Part 1:", count)

	var invalidUpdates []int
	for i := 0; i < len(updates); i++ {
		if contains(validUpdates, i) {
			continue
		}
		invalidUpdates = append(invalidUpdates, i)
	}

	for _, i := range invalidUpdates {
		update := updates[i]
	loop:
		for j, num := range update {
			rule := rules[num]
			for k := 0; k < j; k++ {
				if contains(rule, update[k]) {
					pushBack(update, k)
					goto loop
				}
			}
		}
	}

	count = 0
	for _, i := range invalidUpdates {
		count += middle(updates[i])
	}
	fmt.Println("Part 2:", count)
}

func pushBack(s []int, i int) {
	tmp := s[i]
	copy(s[i:], s[i+1:])
	s[len(s)-1] = tmp
}

func middle(s []int) int {
	l := len(s)
	if l%2 == 0 {
		return s[l/2-1]
	}
	return s[l/2]
}

func contains(rule []int, i int) bool {
	for _, r := range rule {
		if r == i {
			return true
		}
	}
	return false
}
