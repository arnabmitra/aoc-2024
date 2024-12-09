package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/05/input-example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var rules [][2]int
	var sequences [][]int
	// Collect all instructions in order
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				num1, err1 := strconv.Atoi(parts[0])
				num2, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					rules = append(rules, [2]int{num1, num2})
				}
			}
		}
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var seq []int
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err == nil {
					seq = append(seq, num)
				}
			}
			sequences = append(sequences, seq)
		}
	}

	var middleElementsSum int

	// Process each sequence individually
	for _, seq := range sequences {
		// Filter rules relevant to the current sequence
		var filteredRules [][2]int
		seqSet := make(map[int]bool)
		for _, num := range seq {
			seqSet[num] = true
		}
		//filter rules
		for _, rule := range rules {
			if seqSet[rule[0]] && seqSet[rule[1]] {
				filteredRules = append(filteredRules, rule)
			}
		}

		// Build the page order for the current sequence
		order, err := buildPageOrder(filteredRules)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Create a map to store the position of each page in the topological order
		position := make(map[int]int)
		for i, page := range order {
			position[page] = i
		}

		// Validate the current sequence
		if !validateSequence(seq, position) {
			// fix the sequence to be valid
			// find the first element that is not in the correct position
			// find the element that should be in that position
			// swap them
			// validate the sequence again
			// if it is valid, add it to the valid sequences
			// if not, continue*/
			// Fix the sequence to be valid
			for i := 0; i < len(seq); i++ {
				if position[seq[i]] != i {
					// Find the element that should be in this position
					should_be := order[i]
					// Find the position of the element that should be in this position
					should_be_position := findIndex(seq, should_be)
					// Swap the two elements
					swap(seq, i, should_be_position)
				}
			}
			// Take the middle element of the valid sequence
			middle := seq[len(seq)/2]
			middleElementsSum += middle
			fmt.Printf("Middle element of sequence %v is %d\n", seq, middle)
		}
	}

	fmt.Println("Sum of middle elements of all valid sequences:", middleElementsSum)
}

func swap(seq []int, i, j int) {
	temp := seq[i]
	seq[i] = seq[j]
	seq[j] = temp
}

func findIndex(seq []int, should_be int) int {
	for i, val := range seq {
		if val == should_be {
			return i
		}
	}
	return -1 // return -1 if the element is not found
}

// part 2

func main_1() {
	file, err := os.Open("2024/05/input-example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var rules [][2]int
	var sequences [][]int
	// Collect all instructions in order
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				num1, err1 := strconv.Atoi(parts[0])
				num2, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					rules = append(rules, [2]int{num1, num2})
				}
			}
		}
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var seq []int
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err == nil {
					seq = append(seq, num)
				}
			}
			sequences = append(sequences, seq)
		}
	}

	var validSequences [][]int
	var middleElementsSum int

	// Process each sequence individually
	for _, seq := range sequences {
		// Filter rules relevant to the current sequence
		var filteredRules [][2]int
		seqSet := make(map[int]bool)
		for _, num := range seq {
			seqSet[num] = true
		}
		//filter rules
		for _, rule := range rules {
			if seqSet[rule[0]] && seqSet[rule[1]] {
				filteredRules = append(filteredRules, rule)
			}
		}

		// Build the page order for the current sequence
		order, err := buildPageOrder(filteredRules)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Create a map to store the position of each page in the topological order
		position := make(map[int]int)
		for i, page := range order {
			position[page] = i
		}

		// Validate the current sequence
		if validateSequence(seq, position) {
			validSequences = append(validSequences, seq)
			// Take the middle element of the valid sequence
			middle := seq[len(seq)/2]
			middleElementsSum += middle
			fmt.Printf("Middle element of sequence %v is %d\n", seq, middle)
		}
	}

	fmt.Println("Sum of middle elements of all valid sequences:", middleElementsSum)
}

func buildPageOrder(rules [][2]int) ([]int, error) {
	// Step 1: Create the graph
	graph := make(map[int][]int)
	for _, rule := range rules {
		from, to := rule[0], rule[1]
		graph[from] = append(graph[from], to)
	}

	// Step 2: Perform partial ordering using DFS
	visited := make(map[int]bool)
	var order []int
	for node := range graph {
		if !visited[node] {
			dfs(node, graph, visited, &order)
		}
	}

	// Reverse the order to get the correct topological sorting
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	return order, nil
}

func dfs(node int, graph map[int][]int, visited map[int]bool, order *[]int) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited, order)
		}
	}
	*order = append(*order, node)
}

func validateSequence(seq []int, position map[int]int) bool {
	for i := 1; i < len(seq); i++ {
		if position[seq[i-1]] > position[seq[i]] {
			return false
		}
	}
	return true
}
