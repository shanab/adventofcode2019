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

func FindInputs(program []int, result int) (noun int, verb int) {
	np := make([]int, len(program))
	copy(np, program)

	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			np[1], np[2] = n, v
			o := RunProgram(np)[0]
			if o == result {
				return n, v
			}
		}
	}
	return -1, -1
}

const magicResult = 19690720

func main() {
	program, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	noun, verb := FindInputs(program, magicResult)
	if noun == -1 && verb == -1 {
		log.Fatal("Failed to find noun and verb that match")
	}
	fmt.Println("Result:", 100*noun+verb)
}
