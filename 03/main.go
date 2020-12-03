package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	free = '.'
	tree = '#'
)

type grid struct {
	raw           []byte
	width, height int
}

func newGrid(raw []byte) grid {
	var g grid
	g.raw = bytes.ReplaceAll(raw, []byte{'\n'}, []byte{})
	g.width = bytes.IndexByte(raw, '\n')
	g.height = len(g.raw) / g.width
	return g
}

func (g grid) at(x, y int) byte {
	return g.raw[x%g.width+y*g.width]
}

// traverse the grid and count the number of trees.
func (g grid) traverse(dx, dy int) int {
	var num int
	for x, y := 0, 0; y < g.height; x, y = x+dx, y+dy {
		if g.at(x, y) != tree {
			continue
		}
		num++
	}
	return num
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	g := newGrid(raw)
	fmt.Println(part1(g))
	fmt.Println(part2(g))
	return nil
}

func part1(g grid) int {
	return g.traverse(3, 1)
}

func part2(g grid) int {
	num := 1
	num *= g.traverse(1, 1)
	num *= g.traverse(3, 1)
	num *= g.traverse(5, 1)
	num *= g.traverse(7, 1)
	num *= g.traverse(1, 2)
	return num
}
