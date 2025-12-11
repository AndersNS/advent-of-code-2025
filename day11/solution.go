// Package day11
package day11

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc-2025/helpers"
)

const debugMode = false

var l = func() *log.Logger {
	if debugMode {
		return log.New(os.Stdout, "", 0)
	}
	return log.New(io.Discard, "", 0)
}()

type Node struct {
	Key      string
	Children []*Node
}

type Tree struct {
	Nodes map[string]*Node
}

func Newtree(lines []string) *Tree {
	tree := &Tree{Nodes: make(map[string]*Node)}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		key := parts[0]
		childKeys := strings.Fields(parts[1])

		node := tree.GetOrCreate(key)
		for _, child := range childKeys {
			child := tree.GetOrCreate(child)
			node.Children = append(node.Children, child)
		}
	}
	return tree
}

func (t *Tree) GetOrCreate(key string) *Node {
	if node, exists := t.Nodes[key]; exists {
		return node
	}

	node := &Node{Key: key, Children: []*Node{}}
	t.Nodes[key] = node
	return node
}

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	tree := Newtree(lines)
	seen := make(map[string]int64)

	a := getOut(tree, "you", seen)

	l.Printf("Res %d `n", a)

	return strconv.FormatInt(a, 10), nil
}

func getOut(tree *Tree, next string, seen map[string]int64) int64 {
	// base case
	if next == "out" {
		return 1
	}

	// Check if we have been here before
	if val, exists := seen[next]; exists {
		return val
	}

	var paths int64

	if node, exists := tree.Nodes[next]; exists {
		for _, child := range node.Children {
			paths += getOut(tree, child.Key, seen)
		}
	}

	seen[next] = paths
	return paths
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	tree := Newtree(lines)
	memo := make(map[string]int64)

	count := getOut2(tree, "svr", false, false, memo)

	return strconv.FormatInt(count, 10), nil
}

func getOut2(tree *Tree, node string, hasDac, hasFft bool, memo map[string]int64) int64 {
	if node == "dac" {
		hasDac = true
	}
	if node == "fft" {
		hasFft = true
	}

	if node == "out" {
		if hasDac && hasFft {
			return 1
		}
		return 0
	}

	key := fmt.Sprintf("%s:%t:%t", node, hasDac, hasFft)
	if cached, exists := memo[key]; exists {
		return cached
	}

	var paths int64
	if treeNode, exists := tree.Nodes[node]; exists {
		for _, child := range treeNode.Children {
			paths += getOut2(tree, child.Key, hasDac, hasFft, memo)
		}
	}

	memo[key] = paths
	return paths
}
