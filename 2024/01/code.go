package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		// read in input-example.txt
		// input is a string containing the contents of the file
		file, err := os.Open("2024/01/input-example.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return 0
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var col1, col2 []int
		for scanner.Scan() {
			line := scanner.Text()
			columns := strings.Fields(line)
			if len(columns) == 2 {
				val1, _ := strconv.Atoi(columns[0])
				val2, _ := strconv.Atoi(columns[1])
				col1 = append(col1, val1)
				col2 = append(col2, val2)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
		// sort col1 and col2
		sort.Ints(col1)
		sort.Ints(col2)
		fmt.Println("the length of col1 is", len(col1))
		fmt.Println("the length of col2 is", len(col2))
		// find the difference between same index on col1 and col2
		// Calculate the total distance
		totalSimilarity := 0
		for i := 0; i < len(col1); i++ {
			find := col1[i]
			distance := find * findOccurances(col2, find)
			totalSimilarity += distance
		}

		fmt.Println("Total Similarity:", totalSimilarity)
		return totalSimilarity
	}
	// solve part 1 here

	// read in input-example.txt
	// input is a string containing the contents of the file
	file, err := os.Open("2024/01/input-example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var col1, col2 []int
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) == 2 {
			val1, _ := strconv.Atoi(columns[0])
			val2, _ := strconv.Atoi(columns[1])
			col1 = append(col1, val1)
			col2 = append(col2, val2)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	// sort col1 and col2
	sort.Ints(col1)
	sort.Ints(col2)
	fmt.Println("the length of col1 is", len(col1))
	fmt.Println("the length of col2 is", len(col2))
	// find the difference between same index on col1 and col2
	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(col1); i++ {
		distance := abs(col1[i] - col2[i])
		totalDistance += distance
	}

	fmt.Println("Total Distance:", totalDistance)
	return totalDistance
}

func findOccurances(col2 []int, find int) int {
	count := 0
	for _, val := range col2 {
		if val == find {
			count++
		}
	}
	return count
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
