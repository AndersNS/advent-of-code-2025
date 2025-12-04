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
	"aoc-2025/day03"
	"aoc-2025/day04"
)

type SolutionFunc func(string) (string, error)

func main() {
	day := flag.Int("day", 1, "Advent of Code day (1-12)")
	part := flag.Int("part", 1, "Part number (1 or 2)")
	benchmark := flag.Bool("b", false, "Run benchmark (20 iterations)")
	flag.Parse()

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
		3: {
			1: day03.Part1,
			2: day03.Part2,
		},
		4: {
			1: day04.Part1,
			2: day04.Part2,
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

	if *benchmark {
		runBenchmark(solution, inputFile)
	} else {
		start := time.Now()
		res, err := solution(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Result: %s\n", res)
		fmt.Printf("\nCompleted in %v\n", time.Since(start))
	}
}

func runBenchmark(solution SolutionFunc, inputFile string) {
	const warmupRuns = 10
	const iterations = 40

	// Warmup phase
	fmt.Printf("Warming up (%d runs)...\n", warmupRuns)
	for i := 0; i < warmupRuns; i++ {
		_, err := solution(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during warmup: %v\n", err)
			os.Exit(1)
		}
	}

	times := make([]time.Duration, iterations)

	fmt.Printf("\nRunning %d iterations...\n", iterations)
	fmt.Println("---")

	for i := 0; i < iterations; i++ {
		start := time.Now()
		_, err := solution(inputFile)
		elapsed := time.Since(start)
		times[i] = elapsed

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on iteration %d: %v\n", i+1, err)
			os.Exit(1)
		}

		fmt.Printf("Run %2d: %.3fms\n", i+1, float64(elapsed.Microseconds())/1000.0)
	}

	// Calculate statistics
	best := times[0]
	worst := times[0]
	var total time.Duration

	for _, t := range times {
		total += t
		if t < best {
			best = t
		}
		if t > worst {
			worst = t
		}
	}

	average := total / iterations

	fmt.Println("---")
	fmt.Printf("Best:    %.3fms\n", float64(best.Microseconds())/1000.0)
	fmt.Printf("Worst:   %.3fms\n", float64(worst.Microseconds())/1000.0)
	fmt.Printf("Average: %.3fms\n", float64(average.Microseconds())/1000.0)
}
