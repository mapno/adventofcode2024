package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day07/input-0")
	s := bufio.NewScanner(f)

	var sum1, sum2 uint64
	for s.Scan() {
		l := s.Text()
		var (
			targetSum uint64
			numsStr   string
		)
		fmt.Sscanf(l, "%d:", &targetSum)
		numsStr = l[len(l)-len(l)+len(fmt.Sprintf("%d:", targetSum))+1:]
		nums := make([]uint64, 0)
		for _, n := range strings.Split(numsStr, " ") {
			num, _ := strconv.ParseUint(n, 10, 64)
			nums = append(nums, num)
		}
		if matchTarget(targetSum, nums, 1) {
			sum1 += targetSum
		}
		if matchTarget(targetSum, nums, 2) {
			sum2 += targetSum
		}

	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}

func matchTarget(t uint64, nums []uint64, part int) bool {
	ops := make([][]int, 0) // 0 = +, 1 = *, 2 = newNums
	// generate all possible combinations of operations
	generateCombinations(len(nums)-1, []int{}, &ops, part)

	for _, op := range ops {
		sum := nums[0]
		for i, n := range nums[1:] {
			if op[i] == 0 {
				sum += n
			} else if op[i] == 1 {
				sum *= n
			} else {
				sum = concat(sum, n)
			}
			if sum > t {
				break
			}
		}
		if sum == t {
			return true
		}
	}
	return false
}

func concat(a, b uint64) uint64 {
	d := countDigits(b)
	return a*pow(10, d) + b
}

func pow(a, b uint64) uint64 {
	result := uint64(1)
	for i := uint64(0); i < b; i++ {
		result *= a
	}
	return result
}

func countDigits(number uint64) uint64 {
	if number == 0 {
		return 1 // Special case for 0
	}
	count := uint64(0)
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func generateCombinations(n int, current []int, result *[][]int, upTo int) {
	if n == 0 {
		// When n is 0, append a copy of current to result
		// because appending current directly would append a reference.
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	for i := 0; i <= upTo; i++ {
		// Recurse with i added to the current combination
		generateCombinations(n-1, append(current, i), result, upTo)
	}
}
