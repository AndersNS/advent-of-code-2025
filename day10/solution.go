// Package day10
package day10

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"aoc-2025/helpers"
)

func Part1(inputFile string) (string, error) {
	// DEBUG MODE: Set to true for detailed logging, false for minimal output
	const debugMode = true

	// Create a logger that either logs to stdout or discards output
	var l *log.Logger
	if debugMode {
		l = log.New(os.Stdout, "", 0) // Log to stdout with no prefix/timestamp
	} else {
		l = log.New(io.Discard, "", 0) // Discard all output
	}

	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	// XOR LOGIC OVERVIEW:
	// - The pattern (e.g., "#.#") represents a target bit state where # = 1, . = 0
	// - Each button has a bit pattern showing which positions it affects
	// - XOR (^) toggles bits: 0 ^ 1 = 1, 1 ^ 1 = 0
	// - Pressing multiple buttons XORs their patterns together
	// - Goal: Find minimum buttons that XOR to exactly match the target pattern
	//
	// Example: Target = 101 (binary)
	//   Button A = 001, Button B = 100
	//   A ^ B = 001 ^ 100 = 101 ✓ (matches target with 2 buttons)

	totalCost := 0
	for lineNum, li := range lines {
		line, _ := parseLine(li)

		l.Printf("\n=== Line %d ===\n", lineNum+1)
		l.Printf("Pattern: %s\n", line.Pattern)
		l.Printf("Target N (binary): %b (decimal: %d)\n", line.N, line.N)
		l.Printf("Available buttons: %d\n", len(line.Buttons))
		for i, btn := range line.Buttons {
			l.Printf("  Button %d: %b (decimal: %d)\n", i, btn, btn)
		}

		numButtons := len(line.Buttons)
		minButtonsNeeded := math.MaxInt64 // Start with impossible value

		// Try all possible combinations of buttons (2^n combinations)
		// This is brute force: test every subset of buttons to see which XORs to target
		//
		// HOW THE LOOP WORKS:
		// We use 'combination' as a bitmask to represent which buttons to press
		// Example with 3 buttons (A, B, C):
		//   combination = 0 (binary: 000) → press no buttons
		//   combination = 1 (binary: 001) → press button 0 (A)
		//   combination = 2 (binary: 010) → press button 1 (B)
		//   combination = 3 (binary: 011) → press buttons 0,1 (A,B)
		//   combination = 4 (binary: 100) → press button 2 (C)
		//   combination = 5 (binary: 101) → press buttons 0,2 (A,C)
		//   combination = 6 (binary: 110) → press buttons 1,2 (B,C)
		//   combination = 7 (binary: 111) → press buttons 0,1,2 (A,B,C)
		totalCombinations := 1 << numButtons
		l.Printf("\nTrying %d combinations...\n", totalCombinations)

		for combination := range totalCombinations {
			// OPTIMIZATION: Count how many buttons this combination uses
			// We can skip this combination if it uses more buttons than our current best
			numButtonsInCombo := 0
			for bitPos := range numButtons {
				if (combination>>bitPos)%2 == 1 {
					numButtonsInCombo++
				}
			}

			// Early termination: skip if this combination uses too many buttons
			if numButtonsInCombo >= minButtonsNeeded {
				l.Printf("\n  Combination %d: Skipping (uses %d buttons, already found solution with %d)\n",
					combination, numButtonsInCombo, minButtonsNeeded)
				continue
			}

			// XOR starts at 0 (all bits off)
			// As we "press" buttons, we XOR their bit patterns into the result
			xorResult := 0
			buttonsPressed := 0
			selectedButtons := []int{}

			l.Printf("\n  Combination %d (binary: %0*b):\n", combination, numButtons, combination)

			// INNER LOOP: Check each bit position in 'combination' to see which buttons are selected
			// The 'combination' number itself is a bitmask: if bit i is set, button i is pressed
			//
			// Example: combination = 5 (binary: 101), numButtons = 3
			//   bitPos = 0: (5 >> 0) % 2 = 5 % 2 = 1 → button 0 is pressed ✓
			//   bitPos = 1: (5 >> 1) % 2 = 2 % 2 = 0 → button 1 is NOT pressed ✗
			//   bitPos = 2: (5 >> 2) % 2 = 1 % 2 = 1 → button 2 is pressed ✓
			for bitPos := range numButtons {
				// Right shift 'combination' by bitPos positions, then check if lowest bit is 1
				// This extracts the bit at position bitPos
				//
				// DETAILED BREAKDOWN:
				// 1. combination >> bitPos : shift right to move bit bitPos to position 0
				// 2. result % 2 : get the lowest bit (0 or 1)
				// 3. == 1 : check if the bit is set
				bitIsSet := (combination >> bitPos) % 2

				if bitIsSet == 1 {
					l.Printf("    Button %d is PRESSED (bit %d is set in combination %d)\n", bitPos, bitPos, combination)
					l.Printf("      Before XOR: xorResult = %b (decimal: %d)\n", xorResult, xorResult)
					l.Printf("      Button %d value: %b (decimal: %d)\n", bitPos, line.Buttons[bitPos], line.Buttons[bitPos])

					// XOR OPERATION: Toggle the bits that this button affects
					// Example: If xorResult = 0101 and button = 0011
					//          then xorResult ^ button = 0101 ^ 0011 = 0110
					//
					// XOR truth table for each bit position:
					//   0 ^ 0 = 0  (both off → stays off)
					//   0 ^ 1 = 1  (one on → turns on)
					//   1 ^ 0 = 1  (one on → stays on)
					//   1 ^ 1 = 0  (both on → cancels out, turns off)
					//
					// XOR properties:
					//   - Same bit: 0^0=0, 1^1=0 (cancels out)
					//   - Different bit: 0^1=1, 1^0=1 (toggles on)
					//   - Associative: (A^B)^C = A^(B^C) so order doesn't matter
					xorResult ^= line.Buttons[bitPos]

					l.Printf("      After XOR:  xorResult = %b (decimal: %d)\n", xorResult, xorResult)

					buttonsPressed++
					selectedButtons = append(selectedButtons, bitPos)
				} else {
					l.Printf("    Button %d is NOT pressed (bit %d is 0 in combination %d)\n", bitPos, bitPos, combination)
				}
			}

			// Check if the XOR of all selected buttons equals our target pattern
			//
			// Example: Target = 111, Buttons: [001, 010, 100]
			//   After button 0: 001 ≠ 111 (not done yet!)
			//   After button 1: 001 ^ 010 = 011 ≠ 111 (not done yet!)
			//   After button 2: 011 ^ 100 = 111 ✓ (NOW we can check!)
			l.Printf("    Final XOR result: %b, Target: %b\n", xorResult, line.N)
			if xorResult == line.N {
				l.Printf("  ✓✓✓ SUCCESS! Combination %d: buttons %v → XOR result = %b (pressed: %d buttons)\n",
					combination, selectedButtons, xorResult, buttonsPressed)
				l.Printf("      This is now our best solution! (previous best: %d buttons)\n", minButtonsNeeded)
				minButtonsNeeded = min(minButtonsNeeded, buttonsPressed)
			} else {
				l.Printf("  ✗ No match (result %b ≠ target %b)\n", xorResult, line.N)
			}
		}

		if minButtonsNeeded <= numButtons {
			l.Printf("→ Minimum buttons needed: %d\n", minButtonsNeeded)
			totalCost += minButtonsNeeded
		} else {
			l.Printf("→ No valid combination found!\n")
		}
		l.Printf("Running total: %d\n", totalCost)
	}

	fmt.Printf("\n=== FINAL RESULT ===\n")
	fmt.Printf("Total cost: %d\n", totalCost)

	return strconv.Itoa(totalCost), nil
}

type ParsedLine struct {
	N       int
	Pattern string
	Buttons []int
}

func parseLine(line string) (ParsedLine, error) {
	result := ParsedLine{}

	// Extract the pattern in square brackets
	// Example: "[#.#]" means positions 0 and 2 should be "on"
	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	if start == -1 || end == -1 {
		return result, fmt.Errorf("no square brackets found")
	}
	result.Pattern = line[start+1 : end]

	// Convert pattern to target bit representation
	// Each '#' at position i sets bit i to 1
	// Example: "#.#" → positions 0,2 → binary 101 → decimal 5
	N := 0
	for i, c := range line[start+1 : end] {
		if c == '#' {
			// Set bit i using left shift: 1 << i
			// Position 0: 1<<0 = 1 (binary: 001)
			// Position 1: 1<<1 = 2 (binary: 010)
			// Position 2: 1<<2 = 4 (binary: 100)
			N += 1 << i
		}
	}

	result.N = N

	remaining := line[end+1:]

	// Ignore curly braces section (used in part 2)
	if idx := strings.Index(remaining, "{"); idx != -1 {
		remaining = remaining[:idx]
	}

	// Parse button groups in parentheses
	// Each button group like "(0,2,3)" means pressing this button toggles positions 0, 2, and 3
	remaining = strings.TrimSpace(remaining)
	for {
		start := strings.Index(remaining, "(")
		if start == -1 {
			break
		}
		end := strings.Index(remaining, ")")
		if end == -1 {
			break
		}

		groupStr := remaining[start+1 : end]
		group := []int{}

		// Convert button group to bit representation
		// Button affecting positions (0,2,3) → bits 0,2,3 set → binary 1101 → decimal 13
		var s int
		if groupStr != "" {
			parts := strings.Split(groupStr, ",")
			for _, part := range parts {
				num, err := strconv.Atoi(strings.TrimSpace(part))
				// Set bit at position 'num'
				// Example: num=3 → 1<<3 = 8 (binary: 1000)
				s += 1 << num
				if err != nil {
					return result, fmt.Errorf("failed to parse number: %v", err)
				}
				group = append(group, num)
			}
		}

		// Store the button's bit pattern (not the individual positions)
		result.Buttons = append(result.Buttons, s)
		remaining = remaining[end+1:]
	}

	return result, nil
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
