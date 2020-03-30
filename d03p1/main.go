package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"

type Vector struct {
	Direction byte
	Magnitude int
}

func ReadInput() ([][]Vector, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var lines [][]Vector
	for scanner.Scan() {
		rawLine := scanner.Text()
		rawVectors := strings.Split(rawLine, ",")
		var line []Vector
		for _, v := range rawVectors {
			direction := v[0]
			magnitude, err := strconv.Atoi(v[1:])
			if err != nil {
				return nil, err
			}
			line = append(line, Vector{direction, magnitude})
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func FindMinIntersectionDistance(l1, l2 []Vector) int {
	grid := make(map[[2]int][2]bool)

	plotLine(grid, 0, l1)
	plotLine(grid, 1, l2)

	minDistance := math.MaxInt64
	for c, b := range grid {
		if c == [2]int{0, 0} {
			continue
		}
		if d := taxicabDistance(c); b[0] && b[1] && d < minDistance {
			minDistance = d
		}
	}
	return minDistance
}

func taxicabDistance(c [2]int) int {
	return abs(c[0]) + abs(c[1])
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func plotLine(grid map[[2]int][2]bool, l int, line []Vector) {
	var c [2]int
	for _, v := range line {
		for i := 0; i < v.Magnitude; i++ {
			switch v.Direction {
			case 'R':
				c[0]++
			case 'L':
				c[0]--
			case 'U':
				c[1]++
			case 'D':
				c[1]--
			}
			b := grid[c]
			b[l] = true
			grid[c] = b
		}
	}
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Minimum intersection distance:", FindMinIntersectionDistance(lines[0], lines[1]))
}
