package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	paths, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(paths))
	fmt.Println(part2(paths, 100))
	return nil
}

type direction string

func input(r io.Reader) ([][]direction, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	paths := make([][]direction, len(rows))
	for i, r := range rows {
		for j := 0; j < len(r); j++ {
			switch r[j] {
			case 'e', 'w':
				paths[i] = append(paths[i], direction(r[j]))
			case 's', 'n':
				paths[i] = append(paths[i], direction(r[j:j+2]))
				j++
			}
		}
	}
	return paths, nil
}

func part1(paths [][]direction) int {
	g := newGrid()
	for _, p := range paths {
		g.flip(p)
	}
	cs := map[color]int{}
	for _, c := range g.cells {
		cs[c]++
	}
	return cs[black]
}

func part2(paths [][]direction, days int) int {
	g := newGrid()
	for _, p := range paths {
		g.flip(p)
	}
	a := newAutomata(g)
	for i := 0; i < days; i++ {
		a.step()
	}
	cs := map[color]int{}
	for _, c := range a.grid.cells {
		cs[c]++
	}
	return cs[black]
}
