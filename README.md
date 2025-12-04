# Advent of Code 2025

## Usage

Run a specific day and part:

```bash
go run main.go -day 1 -part 1
```

Or test a solution:

```bash
cd day01
go test
```

## Creating a New Day

Copy the `day01` folder structure and update the package name:

```bash
cp -r day03 day04
find day04 -type f -name "*.go" -exec sed -i 's/day03/day04/g' {} +
```

## Hyperfine benchmark

```
go build -o build/aoc main.go
hyperfine --warmup 3 './build/aoc -day 1 -part 2'

```
