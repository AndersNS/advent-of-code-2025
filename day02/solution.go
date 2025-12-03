// Package day02
package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(inputFile string) (string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	parts := strings.Split(strings.TrimSpace(string(data)), ",")

	acc := int64(0)

	for _, part := range parts {
		a := strings.Split(part, "-")
		start, _ := strconv.ParseInt(a[0], 10, 64)
		end, _ := strconv.ParseInt(a[1], 10, 64)
		for id := start; id <= end; id += 1 {
			idStr := strconv.FormatInt(id, 10)
			if len(idStr)%2 != 0 {
				continue
			}
			first := idStr[:len(idStr)/2]
			second := idStr[len(idStr)/2:]

			if second == first {
				acc += id
			}
		}
	}

	return strconv.FormatInt(acc, 10), nil
}

func Part2(inputFile string) (string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	parts := strings.Split(strings.TrimSpace(string(data)), ",")

	acc := int64(0)

	for _, part := range parts {
		a := strings.Split(part, "-")
		start, _ := strconv.ParseInt(a[0], 10, 64)
		end, _ := strconv.ParseInt(a[1], 10, 64)
		for id := start; id <= end; id += 1 {
			idStr := strconv.FormatInt(id, 10)

		patternloop:
			for i := 1; i <= len(idStr)/2; i++ {
				if len(idStr)%i != 0 {
					continue
				}
				pattern := idStr[:i]
				for j := i; j < len(idStr); j++ {
					if idStr[j] != pattern[j%i] {
						continue patternloop
					}
				}
				acc += id
				break
			}
		}
	}

	return strconv.FormatInt(acc, 10), nil
}
