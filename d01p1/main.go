package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename = "input.txt"

func ReadInput() ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var masses []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		masses = append(masses, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return masses, nil
}

func GetFuelSum(masses []int) int {
	var fuelSum int
	for _, m := range masses {
		fuelSum += m/3 - 2
	}
	return fuelSum
}

func main() {
	masses, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FUEL SUM:")
	fmt.Println(GetFuelSum(masses))
}
