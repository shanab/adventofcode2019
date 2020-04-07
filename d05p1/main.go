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

func RunProgram(program []int, input int) []int {
	np := make([]int, len(program))
	copy(np, program)

	var outputs []int

	var i int
	for {
		longCode := fmt.Sprintf("%05d", np[i])
		modes := longCode[:3]
		shortCode := longCode[3:]

		switch shortCode {
		case "01":
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			r := np[i+3]
			np[r] = o1 + o2
			i += 4
		case "02":
			o1 := getOperand(np, i+1, modes[2])
			o2 := getOperand(np, i+2, modes[1])
			r := np[i+3]
			np[r] = o1 * o2
			i += 4
		case "03":
			r := np[i+1]
			np[r] = input
			i += 2
		case "04":
			o := getOperand(np, i+1, modes[2])
			outputs = append(outputs, o)
			i += 2
		case "99":
			return outputs
		}
	}
	return outputs
}

func getOperand(program []int, pos int, mode byte) int {
	if mode == '1' {
		r := program[pos]
		return program[r]
	}
	return program[pos]
}

const magicInput = 1

func main() {
	program, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	outputs := RunProgram(program, magicInput)
	fmt.Println("Outputs:")
	fmt.Println(outputs)
	fmt.Println("Diagnostic Code:")
	fmt.Println(outputs[len(outputs)-1])
}
