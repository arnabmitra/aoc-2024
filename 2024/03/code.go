package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		file, err := os.Open("2024/03/input-example.txt")
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

		// Regex pattern to match all instruction types in order
		pattern := `mul\(\d+,\d+\)|don't\(\)|do\(\)`
		re := regexp.MustCompile(pattern)

		scanner := bufio.NewScanner(file)
		instructions := []string{}

		// Collect all instructions in order
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindAllString(line, -1)
			instructions = append(instructions, matches...)
		}

		// Process instructions in order
		sum := 0
		enabled := true
		for _, instruction := range instructions {
			switch {
			case strings.HasPrefix(instruction, "don't"):
				enabled = false
			case strings.HasPrefix(instruction, "do"):
				enabled = true
			case strings.HasPrefix(instruction, "mul"):
				if enabled {
					sum += multiply(instruction)
				}
			}
		}

		fmt.Println("Sum of all enabled multiplications:", sum)
		return sum
	}
	// solve part 1 here
	file, err := os.Open("2024/03/input-example.txt")
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

	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		// Print the matches

		for _, match := range matches {
			fmt.Println(match)
			sum = sum + multiply(match)
		}
	}
	println("Sum of all matches:", sum)
	return sum
}

func multiply(match string) int {
	num1, num2 := parseNumbers(match)
	return num1 * num2
}

func parseNumbers(input string) (int, int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 3 {
		return 0, 0 // or handle the error as needed
	}
	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])
	return num1, num2
}
