package main

import (
	"bufio"
	"fmt"
	"os"
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
		puzzle = getRuneFromFile()
		rows = len(puzzle)
		cols = len(puzzle[0])

		count := 0
		// Can never start at outer layer, so move bounds in by 1 for efficiency
		for i := 1; i < rows-1; i++ {
			for j := 1; j < cols-1; j++ {
				count += CountCross(i, j)
			}
		}
		println(count)
		return count
	}
	grid := getRuneFromFile()
	count := findWord(grid, "XMAS")
	fmt.Println(count)
	return count
}

const (
	STRAIGHT = 0
	RIGHT    = 1
	LEFT     = -1
	UP       = -1
	DOWN     = 1
)

var puzzle [][]rune
var rows, cols int

func CountCross(row, col int) int {
	count := CountDirection(row+LEFT, col+UP, RIGHT, DOWN, "MAS") +
		CountDirection(row+LEFT, col+DOWN, RIGHT, UP, "MAS") +
		CountDirection(row+RIGHT, col+UP, LEFT, DOWN, "MAS") +
		CountDirection(row+RIGHT, col+DOWN, LEFT, UP, "MAS")

	if count == 2 {
		return 1
	}
	return 0
}

func CountDirection(row, col, dirRow, dirCol int, keyword ...string) int {
	keywordStr := "XMAS"
	if len(keyword) > 0 {
		keywordStr = keyword[0]
	}

	for i := 0; i < len(keywordStr); i++ {
		// Bounds checks
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return 0
		}

		// Check for the keyword
		if puzzle[row][col] != rune(keywordStr[i]) {
			return 0
		}

		row += dirRow
		col += dirCol
	}

	return 1
}

func getRuneFromFile() [][]rune {
	// solve part 1 here
	file, err := os.Open("2024/04/input-example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	// Create the 2D rune grid
	var grid [][]rune
	// Collect all instructions in order
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid
}

func findWord(grid [][]rune, word string) int {
	wordLength := len(word)
	rows := len(grid)
	cols := len(grid[0])
	wordRunes := []rune(word)

	// Directions: right, left, down, up, diagonals
	directions := [][]int{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // down
		{-1, 0},  // up
		{1, 1},   // diagonal down-right
		{-1, -1}, // diagonal up-left
		{1, -1},  // diagonal down-left
		{-1, 1},  // diagonal up-right
	}

	occurrences := 0

	// Function to check if the word exists in a given direction
	isWordInDirection := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLength; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != wordRunes[i] {
				return false
			}
		}
		return true
	}

	// Iterate over every cell in the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check each direction from this cell
			for _, dir := range directions {
				if isWordInDirection(i, j, dir[0], dir[1]) {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func createRuneGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
