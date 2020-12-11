package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	g, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(g.clone()))
	fmt.Println(part2(g.clone()))
	return nil
}

func input(r io.Reader) (*grid, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return newGrid(data), nil
}

func part1(g *grid) int {
	fn := func(cell byte, adj []byte) (byte, bool) {
		if cell == cellEmpty && bytes.Count(adj, []byte{cellOccupied}) == 0 {
			return cellOccupied, true
		}
		if cell == cellOccupied && bytes.Count(adj, []byte{cellOccupied}) >= 4 {
			return cellEmpty, true
		}
		return cellFloor, false
	}
	for g.evolve(fn, distShortsighted) {
	}
	return bytes.Count(g.data, []byte{cellOccupied})
}

func part2(g *grid) int {
	fn := func(cell byte, adj []byte) (byte, bool) {
		if cell == cellEmpty && bytes.Count(adj, []byte{cellOccupied}) == 0 {
			return cellOccupied, true
		}
		if cell == cellOccupied && bytes.Count(adj, []byte{cellOccupied}) >= 5 {
			return cellEmpty, true
		}
		return cellFloor, false
	}
	for g.evolve(fn, distInfinity) {
	}
	return bytes.Count(g.data, []byte{cellOccupied})
}
