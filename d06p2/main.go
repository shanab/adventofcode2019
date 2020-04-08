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
K)L
K)YOU
I)SAN`

type Node struct {
	Key      string
	Edges    []*Node
	Distance int
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
		nodes[pk].Edges = append(nodes[pk].Edges, nodes[ck])
		nodes[ck].Edges = append(nodes[ck].Edges, nodes[pk])
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return nodes, nil
}

func PathDistance(node *Node, dst string) int {
	visited := map[string]bool{node.Key: true}
	queue := []*Node{node}

	var n *Node
	for len(queue) != 0 {
		n, queue = queue[len(queue)-1], queue[:len(queue)-1]
		visited[n.Key] = true
		if n.Key == dst {
			// Due to wording of the problem,
			// distance returned should be subtracted by 2
			return n.Distance - 2
		}
		for _, e := range n.Edges {
			if visited[e.Key] {
				continue
			}
			e.Distance = n.Distance + 1
			queue = append(queue, e)
		}
	}
	return -1
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
	fmt.Println("Distance between YOU and SAN:")
	fmt.Println(PathDistance(nodes["YOU"], "SAN"))
}
