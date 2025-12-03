// Description: Main entry point for Advent of Code 2025 solutions.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"aoc-2025/day01"
	"aoc-2025/day02"
)

type SolutionFunc func(string) (string, error)

func main() {
	day := flag.Int("day", 1, "Advent of Code day (1-12)")
	part := flag.Int("part", 1, "Part number (1 or 2)")
	flag.Parse()

	if *day < 1 || *day > 12 {
		fmt.Println("Day must be between 1 and 25")
		os.Exit(1)
	}

	if *part < 1 || *part > 2 {
		fmt.Println("Part must be 1 or 2")
		os.Exit(1)
	}

	fmt.Printf("Running Day %d, Part %d\n", *day, *part)
	fmt.Println("---")
	solutions := map[int]map[int]SolutionFunc{
		1: {
			1: day01.Part1,
			2: day01.Part2,
		},
		2: {
			1: day02.Part1,
			2: day02.Part2,
		},
	}
	inputFile := filepath.Join(fmt.Sprintf("day%02d", *day), "input.txt")

	dayMap, ok := solutions[*day]
	if !ok {
		fmt.Printf("Day %d not yet implemented\n", *day)
		os.Exit(1)
	}
	solution, ok := dayMap[*part]
	if !ok {
		fmt.Printf("Day %d Part %d not found\n", *day, *part)
		os.Exit(1)
	}
	start := time.Now()
	res, err := solution(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Result: %s\n", res)

	fmt.Printf("\nCompleted in %v\n", time.Since(start))
}
