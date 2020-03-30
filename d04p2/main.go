package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const filename = "input.txt"

func ReadInput() (int, int, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, 0, err
	}
	raw := strings.TrimSpace(string(dat))
	rawNums := strings.Split(raw, "-")
	if len(rawNums) != 2 {
		return 0, 0, errors.New("Invalid input: expected 2 numbers")
	}
	n1, err := strconv.Atoi(rawNums[0])
	if err != nil {
		return 0, 0, errors.New("Invalid input: first number invalid")
	}
	n2, err := strconv.Atoi(rawNums[1])
	if err != nil {
		return 0, 0, errors.New("Invalid input: second number invalid")
	}
	return n1, n2, nil
}

func CountValidPasswords(start, end int) int {
	var count int
	for i := start; i <= end; i++ {
		if isValidPassword(i) {
			count++
		}
	}
	return count
}

func isValidPassword(n int) bool {
	s := strconv.Itoa(n)
	c := s[0]
	doubleDigit := s[0]
	doubleCount := 1
	var foundDouble bool

	for i := 1; i < len(s); i++ {
		if s[i] < c {
			return false
		}

		if !foundDouble {
			if s[i] == doubleDigit {
				doubleCount++
			} else {
				foundDouble = (doubleCount == 2)
				doubleDigit = s[i]
				doubleCount = 1
			}
		}

		c = s[i]
	}

	return foundDouble || doubleCount == 2
}

func main() {
	start, end, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of valid passwords:", CountValidPasswords(start, end))
}
