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
		fuel := getFuel(m)
		fuelSum += fuel + getExtraFuel(fuel)
	}
	return fuelSum
}

func getExtraFuel(fuel int) int {
	var total int
	f := fuel
	for {
		m := getFuel(f)
		if m < 0 {
			break
		}
		total += m
		f = m
	}
	return total
}

func getFuel(mass int) int {
	return mass/3 - 2
}

func main() {
	masses, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fuelSum := GetFuelSum(masses)
	fmt.Println("Fuel sum with extra fuel:", fuelSum)
}
