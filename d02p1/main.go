package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const filename = "input.txt"

func ReadInput() ([]int, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var numbers []int
	raw := strings.TrimSpace(string(dat))
	for _, s := range strings.Split(raw, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func RunProgram(program []int) []int {
	np := make([]int, len(program))
	copy(np, program)

Loop:
	for i := 0; ; i += 4 {
		opcode := np[i]
		r1, r2, r3 := np[i+1], np[i+2], np[i+3]
		switch opcode {
		case 1:
			np[r3] = np[r1] + np[r2]
		case 2:
			np[r3] = np[r1] * np[r2]
		case 99:
			break Loop
		}
	}
	return np
}

func main() {
	program, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	// Adjust input, just because
	program[1], program[2] = 12, 2

	newProgram := RunProgram(program)
	fmt.Println("Value at position 0:", newProgram[0])
	fmt.Println("program:", program)
	fmt.Println("nprogra:", newProgram)
}
