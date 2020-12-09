package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	preamble := flag.Int("preamble", 25, "Number of items in the preamble")
	flag.Parse()
	if err := run(*preamble); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(preamble int) error {
	is, err := input(os.Stdin)
	if err != nil {
		return err
	}
	a := part1(is, preamble)
	fmt.Println(a)
	b := part2(is, a)
	fmt.Println(b)
	return nil
}

type system struct {
	input    []int
	ptr      int
	preamble int
	err      error
}

func newSystem(input []int, preamble int) *system {
	return &system{input: input, preamble: preamble, ptr: preamble}
}

func (s *system) next() bool {
	preamble := map[int]interface{}{}
	for _, v := range s.input[s.ptr-s.preamble : s.ptr] {
		preamble[v] = nil
	}
	v := s.input[s.ptr]
	for w := range preamble {
		if v%2 == 0 && v/2 == w {
			continue
		}
		if _, ok := preamble[v-w]; ok {
			s.ptr++
			return true
		}
	}
	return false
}

func (s *system) value() int { return s.input[s.ptr] }

func input(r io.Reader) ([]int, error) {
	var is []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		is = append(is, i)
	}
	return is, s.Err()
}

func part1(input []int, preamble int) int {
	s := newSystem(input, preamble)
	for s.next() {
	}
	return s.value()
}

func part2(input []int, weakness int) int {
	for i := range input {
		var acc int
		var min, max int
		for _, j := range input[i:] {
			acc += j
			if j < min || min == 0 {
				min = j
			}
			if j > max || min == 0 {
				max = j
			}
			if acc > weakness {
				break
			}
			if acc == weakness {
				return min + max
			}
		}
	}
	return 0
}
