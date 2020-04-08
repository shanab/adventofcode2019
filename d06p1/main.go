package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFile = "input.txt"

const testInput = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

type Node struct {
	Key    string
	Parent *Node
}

func (n Node) Orbits() int {
	if n.Parent == nil {
		return 0
	}
	return 1 + n.Parent.Orbits()
}

func ReadInput(r io.Reader) (map[string]*Node, error) {
	sc := bufio.NewScanner(r)
	nodes := make(map[string]*Node)
	for sc.Scan() {
		line := sc.Text()
		objects := strings.Split(line, ")")
		pk, ck := objects[0], objects[1]

		if nodes[pk] == nil {
			nodes[pk] = &Node{Key: pk}
		}
		if nodes[ck] == nil {
			nodes[ck] = &Node{Key: ck}
		}
		nodes[ck].Parent = nodes[pk]
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return nodes, nil
}

func CountOrbits(nodes map[string]*Node) int {
	var result int
	for _, n := range nodes {
		result += n.Orbits()
	}
	return result
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	nodes, err := ReadInput(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of direct and indirect orbits:")
	fmt.Println(CountOrbits(nodes))
}
