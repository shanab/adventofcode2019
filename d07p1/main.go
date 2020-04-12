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

func RunProgram(program []int, inputs []int) []int {
	np := make([]int, len(program))
	copy(np, program)

	var im int
	var outputs []int

	var i int
	for {
		longCode := fmt.Sprintf("%05d", np[i])
		modes := longCode[:3]
		shortCode := longCode[3:]

		switch shortCode {
		case "01": // add
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			r := np[i+3]
			np[r] = o1 + o2
			i += 4
		case "02": // multiply
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			r := np[i+3]
			np[r] = o1 * o2
			i += 4
		case "03": // input
			r := np[i+1]
			np[r] = inputs[im]
			im++
			i += 2
		case "04": // output
			o := getOperand(np, i+1, modes[2])
			outputs = append(outputs, o)
			i += 2
		case "05": // Jump If True
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			if o1 != 0 {
				i = o2
				continue
			}
			i += 3
		case "06": // jump if false
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			if o1 == 0 {
				i = o2
				continue
			}
			i += 3
		case "07": // store if less than
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			r := np[i+3]
			if o1 < o2 {
				np[r] = 1
			} else {
				np[r] = 0
			}
			i += 4
		case "08": // store if equal
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			r := np[i+3]
			if o1 == o2 {
				np[r] = 1
			} else {
				np[r] = 0
			}
			i += 4
		case "99": // halt
			return outputs
		}
	}
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
	maxSignal := 0
	for _, perm := range perms {
		// Reset input every time a new permutation is tried
		in := input
		for _, a := range perm {
			outputs := RunProgram(program, []int{a, in})
			in = outputs[0]
			if outputs[0] > maxSignal {
				maxSignal = outputs[0]
			}
		}
	}
	return maxSignal
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

var settings []int = []int{0, 1, 2, 3, 4}

const input = 0

func main() {
	program, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Max thrusters signal:")
	fmt.Println(MaxThrusterSignal(program, input, settings))
}
