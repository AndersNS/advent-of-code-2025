// Package day08
package day08

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"aoc-2025/helpers"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

// UnionFind Trying to implement unionfind in go, gg
type UnionFind struct {
	parent []int
	rank   []int
}

func (p Point) Distance(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	dz := p.Z - other.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}
	return uf
}

// Find finds the root of the cluster containing point at x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Cluster Just a helper for easier printing
type Cluster struct {
	Root    int
	Members []int
}

func (uf *UnionFind) GetClusters() []Cluster {
	clusterMap := make(map[int][]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		clusterMap[root] = append(clusterMap[root], i)
	}

	// Convert to slice
	clusters := make([]Cluster, 0, len(clusterMap))
	for root, members := range clusterMap {
		clusters = append(clusters, Cluster{Root: root, Members: members})
	}

	// Sort by size (largest first)
	sort.Slice(clusters, func(i, j int) bool {
		return len(clusters[i].Members) > len(clusters[j].Members)
	})

	return clusters
}

// Union Merges the clusters containing x and y by finding their roots and linking them
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

func (uf *UnionFind) CountClusters() int {
	// Find unique roots
	roots := make(map[int]bool)
	for i := 0; i < len(uf.parent); i++ {
		roots[uf.Find(i)] = true
	}
	return len(roots)
}

type PointPair struct {
	Index1   int
	Index2   int
	Distance float64
}

type IncrementalClusterer struct {
	points      []Point
	uf          *UnionFind
	pairs       []PointPair
	step        int
	step2Answer float64 // Yeah
}

func NewIncrementalClusterer(points []Point) *IncrementalClusterer {
	n := len(points)
	ic := &IncrementalClusterer{
		points: points,
		uf:     NewUnionFind(n),
		pairs:  make([]PointPair, 0),
		step:   0,
	}

	// Calculate all pairwise distances
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := points[i].Distance(points[j])
			ic.pairs = append(ic.pairs, PointPair{
				Index1:   i,
				Index2:   j,
				Distance: dist,
			})
		}
	}

	// Sort pairs by distance (closest first)
	sort.Slice(ic.pairs, func(a, b int) bool {
		return ic.pairs[a].Distance < ic.pairs[b].Distance
	})

	return ic
}

type Distance struct {
	Index    int
	Distance float64
}

// Step performs one merge operation (combines the next closest pair)

func (ic *IncrementalClusterer) Step() bool {
	if ic.step >= len(ic.pairs) {
		return false // No more pairs to process
	}

	pair := ic.pairs[ic.step]
	ic.step++

	// Try to merge - returns true if they were in different clusters
	merged := ic.uf.Union(pair.Index1, pair.Index2)

	if merged {
		p1 := ic.points[pair.Index1]
		p2 := ic.points[pair.Index2]
		fmt.Printf("Step %d: Merged point %d (%.1f,%.1f,%.1f) with point %d (%.1f,%.1f,%.1f) [distance: %.2f]\n",
			ic.step, pair.Index1, p1.X, p1.Y, p1.Z, pair.Index2, p2.X, p2.Y, p2.Z, pair.Distance)
		ic.step2Answer = p1.X * p2.X
		fmt.Printf("\tClusters remaining: %d\n\n", ic.uf.CountClusters())
	}

	return true
}

func (ic *IncrementalClusterer) StepN(n int) {
	for i := 0; i < n; i++ {
		if !ic.Step() {
			break
		}
	}
}

func (ic *IncrementalClusterer) OneCluster() {
	for ic.Step() {
	}
}

func (ic *IncrementalClusterer) GetCurrentClusters() []Cluster {
	return ic.uf.GetClusters()
}

func Part1(inputFile string) (string, error) {
	// different parameters for test data and real data
	return part1Internal(inputFile, 1000, 3)
}

func part1Internal(inputFile string, steps int, sum int) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var points []Point

	for _, line := range lines {
		a := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(a[0], 64)
		y, _ := strconv.ParseFloat(a[1], 64)
		z, _ := strconv.ParseFloat(a[2], 64)

		points = append(points, Point{X: x, Y: y, Z: z})
	}

	clusterer := NewIncrementalClusterer(points)
	clusterer.StepN(steps)

	clusters := clusterer.GetCurrentClusters()
	printClusters(points, clusters)
	acc := len(clusters[0].Members)
	for s := 1; s < sum; s++ {
		acc *= len(clusters[s].Members)
	}

	return strconv.FormatInt(int64(acc), 10), nil
}

func printClusters(points []Point, clusters []Cluster) {
	for i, cluster := range clusters {
		fmt.Printf("Cluster %d (root: %d, size: %d):\n", i+1, cluster.Root, len(cluster.Members))
		for _, idx := range cluster.Members {
			p := points[idx]
			fmt.Printf("  Point %d: (%.1f, %.1f, %.1f)\n", idx, p.X, p.Y, p.Z)
		}
		fmt.Println()
	}
	fmt.Printf("Total clusters: %d\n", len(clusters))
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var points []Point

	for _, line := range lines {
		a := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(a[0], 64)
		y, _ := strconv.ParseFloat(a[1], 64)
		z, _ := strconv.ParseFloat(a[2], 64)

		points = append(points, Point{X: x, Y: y, Z: z})
	}

	clusterer := NewIncrementalClusterer(points)
	clusterer.OneCluster()

	// Show current clusters
	fmt.Println("=== Current cluster state ===")
	clusters := clusterer.GetCurrentClusters()
	printClusters(points, clusters)

	return strconv.FormatInt(int64(clusterer.step2Answer), 10), nil
}
