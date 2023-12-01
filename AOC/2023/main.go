package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input []string
	// Day1
	input, _ = readLines("day1.txt")

	fmt.Printf("Day 1 Part 1: %v\n", day1Part1(input))
	fmt.Printf("Day 1 Part 2: %v\n", day1Part2(input))
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	// return the reversed string.
	return string(rns)
}
