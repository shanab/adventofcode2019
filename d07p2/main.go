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

type Computer struct {
	unmodifiedProgram []int
	program           []int
	pc                int
}

func NewComputer(program []int) *Computer {
	c := Computer{
		program:           make([]int, len(program)),
		unmodifiedProgram: make([]int, len(program)),
	}
	copy(c.program, program)
	copy(c.unmodifiedProgram, program)
	return &c
}

func (c *Computer) RunProgram(inputs []int) ([]int, bool) {
	var im int
	var outputs []int
	for {
		longCode := fmt.Sprintf("%05d", c.program[c.pc])
		modes := longCode[:3]
		shortCode := longCode[3:]

		switch shortCode {
		case "01": // add
			o1 := getOperand(c.program, c.pc+1, modes[2])
			o2 := getOperand(c.program, c.pc+2, modes[1])
			r := c.program[c.pc+3]
			c.program[r] = o1 + o2
			c.pc += 4
		case "02": // multiply
			o1 := getOperand(c.program, c.pc+1, modes[2])
			o2 := getOperand(c.program, c.pc+2, modes[1])
			r := c.program[c.pc+3]
			c.program[r] = o1 * o2
			c.pc += 4
		case "03": // input
			if im >= len(inputs) {
				return outputs, false
			}
			r := c.program[c.pc+1]
			c.program[r] = inputs[im]
			im++
			c.pc += 2
		case "04": // output
			o := getOperand(c.program, c.pc+1, modes[2])
			outputs = append(outputs, o)
			c.pc += 2
		case "05": // Jump If True
			o1 := getOperand(c.program, c.pc+1, modes[2])
			o2 := getOperand(c.program, c.pc+2, modes[1])
			if o1 != 0 {
				c.pc = o2
				continue
			}
			c.pc += 3
		case "06": // jump if false
			o1 := getOperand(c.program, c.pc+1, modes[2])
			o2 := getOperand(c.program, c.pc+2, modes[1])
			if o1 == 0 {
				c.pc = o2
				continue
			}
			c.pc += 3
		case "07": // store if less than
			o1 := getOperand(c.program, c.pc+1, modes[2])
			o2 := getOperand(c.program, c.pc+2, modes[1])
			r := c.program[c.pc+3]
			if o1 < o2 {
				c.program[r] = 1
			} else {
				c.program[r] = 0
			}
			c.pc += 4
		case "08": // store if equal
			o1 := getOperand(c.program, c.pc+1, modes[2])
			o2 := getOperand(c.program, c.pc+2, modes[1])
			r := c.program[c.pc+3]
			if o1 == o2 {
				c.program[r] = 1
			} else {
				c.program[r] = 0
			}
			c.pc += 4
		case "99": // halt
			return outputs, true
		}
	}
}

func (c *Computer) Reset() {
	c.program = make([]int, len(c.unmodifiedProgram))
	copy(c.program, c.unmodifiedProgram)
	c.pc = 0
}
func getOperand(program []int, pos int, mode byte) int {
	if mode == '0' {
		r := program[pos]
		return program[r]
	}
	return program[pos]
}

func MaxThrusterSignal(program []int, input int, settings []int) int {
	perms := permutations(settings)
	computers := make([]*Computer, len(settings))
	for i := range settings {
		computers[i] = NewComputer(program)
	}
	maxSignal := 0
	for _, perm := range perms {
		signal := ThrusterSignal(computers, input, perm)
		if signal > maxSignal {
			maxSignal = signal
		}
		for _, c := range computers {
			c.Reset()
		}
	}
	return maxSignal
}

func ThrusterSignal(computers []*Computer, input int, perm []int) int {
	channels := make([][]int, len(perm))
	for i := 0; i < len(perm); i++ {
		channels[i] = append(channels[i], perm[i])
	}
	channels[0] = append(channels[0], input)

	for {
		for i, c := range computers {
			// Run program
			outputs, halted := c.RunProgram(channels[i])
			if halted && i == len(perm)-1 {
				return outputs[len(outputs)-1]
			}

			// Clear already used input
			channels[i] = nil

			// Append output as input to the next amplifier
			j := (i + 1) % len(perm)
			channels[j] = append(channels[j], outputs...)
		}
	}
}

func permutations(a []int) [][]int {
	var results [][]int
	permute(a, 0, func(a []int) {
		results = append(results, a)
	})
	return results
}

func permute(a []int, start int, f func([]int)) {
	if start >= len(a) {
		c := make([]int, len(a))
		copy(c, a)
		f(c)
		return
	}

	for i := start; i < len(a); i++ {
		a[i], a[start] = a[start], a[i]
		permute(a, start+1, f)
		a[start], a[i] = a[i], a[start]
	}
}

var settings []int = []int{5, 6, 7, 8, 9}

const magicInput = 0

func main() {
	program, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Max thrusters signal:")
	fmt.Println(MaxThrusterSignal(program, magicInput, settings))
}
