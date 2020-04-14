package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

const filename = "input.txt"

func ReadInput() ([]int, error) {
	dat, err := ioutil.ReadFile(filename)
	dat = bytes.TrimSpace(dat)
	if err != nil {
		return nil, err
	}
	nums := make([]int, len(dat))
	for i, b := range dat {
		nums[i] = int(b - '0')
	}
	return nums, nil
}

const (
	length = 6
	width  = 25
)

func FindResult(pixels []int) int {
	layerSize := length * width
	layers := len(pixels) / layerSize

	var result int
	fewest := math.MaxInt64
	for i := 0; i < layers; i++ {
		var zeros, ones, twos int
		for j := 0; j < layerSize; j++ {
			k := i*layerSize + j
			switch pixels[k] {
			case 0:
				zeros++
			case 1:
				ones++
			case 2:
				twos++
			}
		}
		if zeros < fewest {
			fewest = zeros
			result = ones * twos
		}
	}

	return result
}

func main() {
	pixels, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:")
	fmt.Println(FindResult(pixels))
}
