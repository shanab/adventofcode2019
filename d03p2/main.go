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

type Point struct {
	Visited bool
	Steps   int
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

func FindMinIntersectionSteps(l1, l2 []Vector) int {
	grid := make(map[[2]int][2]Point)

	plotLine(grid, 0, l1)
	plotLine(grid, 1, l2)

	minSteps := math.MaxInt64
	for c, b := range grid {
		if c == [2]int{0, 0} {
			continue
		}
		if s := b[0].Steps + b[1].Steps; b[0].Visited && b[1].Visited && s < minSteps {
			minSteps = s
		}
	}
	return minSteps
}

func plotLine(grid map[[2]int][2]Point, l int, line []Vector) {
	var c [2]int
	var steps int
	for _, v := range line {
		for i := 0; i < v.Magnitude; i++ {
			steps++

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

			p := grid[c]
			if !p[l].Visited {
				p[l].Visited = true
				p[l].Steps = steps
				grid[c] = p
			}
		}
	}
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Minimum intersection steps:", FindMinIntersectionSteps(lines[0], lines[1]))
}
