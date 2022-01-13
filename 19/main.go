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
	rs, ms, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(rs, ms))
	fmt.Println(part2(rs, ms))
	return nil
}

func input(r io.Reader) (map[string][][]string, []string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	// The first part contains the rules, the second messages.
	parts := strings.Split(string(data), "\n\n")

	// Rules are slices of possible pointers or values.
	rs := map[string][][]string{}
	for _, l := range strings.Split(parts[0], "\n") {
		if l == "" {
			continue
		}
		sides := strings.SplitN(l, ": ", 2)
		left, right := sides[0], sides[1]
		right = strings.ReplaceAll(right, "\"", "")
		paths := strings.Split(right, " | ")
		for _, p := range paths {
			rs[left] = append(rs[left], strings.Fields(p))
		}
	}

	ms := strings.Split(strings.TrimSpace(parts[1]), "\n")

	return rs, ms, nil
}

func part1(rs map[string][][]string, ms []string) int {
	var sum int
	for _, m := range ms {
		for _, n := range match(m, "0", 0, rs) {
			if n == len(m) {
				sum++
			}
		}
	}
	return sum
}

func part2(rs map[string][][]string, ms []string) int {
	var sum int
	rs["8"] = [][]string{{"42"}, {"42", "8"}}
	rs["11"] = [][]string{{"42", "31"}, {"42", "11", "31"}}
	for _, m := range ms {
		for _, n := range match(m, "0", 0, rs) {
			if n == len(m) {
				sum++
			}
		}
	}
	return sum
}

func match(str string, key string, index int, rs map[string][][]string) []int {
	if index >= len(str) {
		return nil
	}

	var ns []int
	for _, r := range rs[key] {
		if len(r) == 1 && r[0][0] >= 'a' && r[0][0] <= 'z' {
			if string(str[index]) != r[0] {
				return nil
			}
			return []int{index + 1}
		}

		ms := []int{index}
		for _, b := range r {
			var os []int
			for _, i := range ms {
				os = append(os, match(str, string(b), i, rs)...)
			}
			ms = os
		}
		ns = append(ns, ms...)
	}
	return ns
}
