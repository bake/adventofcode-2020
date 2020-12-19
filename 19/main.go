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
		n, _ := match(m, "0", rs)
		if n == len(m) {
			sum++
		}
	}
	return sum
}

func match(str string, key string, rs map[string][][]string) (int, bool) {
	var n int
	for _, r := range rs[key] {
		cpy := str
		if len(cpy) >= 1 && len(r) == 1 && r[0][0] >= 'a' && r[0][0] <= 'z' {
			if string(cpy[0]) == r[0] {
				return 1, true
			}
			return 0, false
		}
		var m int
		for _, b := range r {
			o, ok := match(cpy, string(b), rs)
			if !ok {
				break
			}
			m, cpy = m+o, cpy[o:]
		}
		if m > n {
			n = m
		}
	}
	return n, n > 0
}
