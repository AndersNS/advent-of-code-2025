package day00

import (
	"fmt"
	"os"
)

func Part1(inputFile string) (string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	fmt.Println("Day 01, Part 1")
	fmt.Printf("Input length: %d bytes\n", len(data))
	fmt.Println("TODO: Implement solution")

	return "", nil
}

func Part2(inputFile string) (string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	fmt.Println("Day 01, Part 2")
	fmt.Printf("Input length: %d bytes\n", len(data))
	fmt.Println("TODO: Implement solution")

	return "", nil
}
