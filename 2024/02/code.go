package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	run(true, "")
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		file, err := os.Open("2024/02/input-example.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return 0
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("Error closing file:", err)
			}
		}(file)
		scanner := bufio.NewScanner(file)
		safeCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			levels := strings.Fields(line)

			if isSafe(levels) || canBeSafeWithOneRemoval(levels) {
				safeCount++
			}

		}
		fmt.Println("Total safe:", safeCount)
		return safeCount

	}
	// solve part 1 here
	file, err := os.Open("2024/02/input-example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)

		if isSafe(levels) {
			safeCount++
		}

	}
	fmt.Println("Total safe:", safeCount)
	return safeCount
}

func isSafe(levels []string) bool {
	if len(levels) < 2 {
		return false
	}
	first, _ := strconv.Atoi(levels[0])
	increasing := true
	decreasing := true

	for i := 1; i < len(levels); i++ {
		current, _ := strconv.Atoi(levels[i])
		// current 4 , first 10
		diff := abs(current - first)

		if diff < 1 || diff > 3 {
			return false
		}
		if current > first {
			decreasing = false
		} else if current < first {
			increasing = false
		} else {
			return false
		}

		first = current

	}
	return increasing || decreasing
}
func canBeSafeWithOneRemoval(levels []string) bool {
	for i := 0; i < len(levels); i++ {
		newLevels := make([]string, len(levels)-1)
		copy(newLevels, levels[:i])
		copy(newLevels[i:], levels[i+1:])
		fmt.Printf("%v \n", newLevels)
		if isSafe(newLevels) {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
