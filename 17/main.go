package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	gens := flag.Int("generations", 6, "Number of generations")
	flag.Parse()
	if err := run(*gens); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(gens int) error {
	rect, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fn := func(active bool, neighbors int) bool {
		if active && (neighbors == 2 || neighbors == 3) {
			return true
		}
		if !active && (neighbors == 3) {
			return true
		}
		return false
	}
	fmt.Println(part1(newCube(rect), gens, fn))
	fmt.Println(part2(newHypercube(newCube(rect)), gens, fn))
	return nil
}

type evolveFn func(active bool, neighbors int) bool

func part1(c *cube, n int, fn evolveFn) int {
	for i := 0; i < n; i++ {
		c.evolve(fn)
	}
	return len(c.data)
}

func part2(c *hypercube, n int, fn evolveFn) int {
	for i := 0; i < n; i++ {
		c.evolve(fn)
	}
	return len(c.data)
}

func input(r io.Reader) (*rectangle, error) {
	bounds := bounds2d{}
	data := map[point2d]interface{}{}
	s := bufio.NewScanner(r)
	for y := 0; s.Scan(); y++ {
		for x, v := range s.Text() {
			switch v {
			case '.':
			case '#':
				if bounds.max.x < x {
					bounds.max.x = x
				}
				if bounds.max.y < y {
					bounds.max.y = y
				}
				data[point2d{x, y}] = nil
			}
		}
	}
	return &rectangle{bounds, data}, nil
}
