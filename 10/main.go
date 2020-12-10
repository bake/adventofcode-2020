package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	is, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(is))
	fmt.Println(part2(is))
	return nil
}

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
	sort.Ints(is)
	return is, s.Err()
}

func part1(is []int) int {
	diffs := map[int]int{}
	var prev int
	for _, i := range is {
		diffs[i-prev]++
		prev = i
	}
	return diffs[1] * (diffs[3] + 1)
}

// tribonacci returns the nth number in the tribonacci sequence. This could be
// stored in memory but given the puzzle input n won't be any bigger than 4.
func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return a
}

// part2 calculates the number of ways to arrange the adapters. Is this readable
// code? I'd argue that it isn't. But it works.
func part2(is []int) int {
	is = append([]int{0}, is...)
	is = append(is, is[len(is)-1]+3)
	prev, prod := 0, 1
	for i := 1; i < len(is); i++ {
		if is[i]-is[i-1] < 3 {
			continue
		}
		prev, prod = i, prod*tribonacci(i-prev)
	}
	return prod
}
