package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

const filename = "input.txt"

func ReadInput() ([]byte, error) {
	dat, err := ioutil.ReadFile(filename)
	dat = bytes.TrimSpace(dat)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func DecodeImage(pixels []byte) []byte {
	layerSize := length * width

	// Initialize a flat transparent image formed of a long series of pixels
	image := make([]byte, layerSize)
	copy(image, pixels[:layerSize])

	for i := 0; i < len(pixels); i++ {
		j := i % layerSize
		if image[j] == '2' {
			image[j] = pixels[i]
		}
	}

	return image
}

func PrintImage(image []byte) {
	for i := 0; i < len(image); i++ {
		if i%(width) == 0 {
			fmt.Println()
		}
		if image[i] == '1' {
			fmt.Print("1")
		} else {
			fmt.Print(" ")
		}

	}
	fmt.Println()
}

const (
	length = 6
	width  = 25
)

func main() {
	pixels, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	image := DecodeImage(pixels)
	fmt.Println("Image")
	PrintImage(image)
}
