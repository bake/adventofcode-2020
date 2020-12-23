package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	l, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(l.copy(), 100))
	fmt.Println(part2(l.copy(), 10_000_000, 1_000_000))
	return nil
}

func input(r io.Reader) (*list, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var l list
	for _, r := range data {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, err
		}
		l.add(&node{label: v})
	}
	l.next()
	return &l, nil
}

func part1(l *list, moves int) string {
	min, max := l.min().label, l.max().label
	for i := 0; i < moves; i++ {
		current := l.current
		l.next()
		a, b, c := l.take(), l.take(), l.take()
		var dest *node
		for i := current.label - 1; dest == nil; i-- {
			if i < min {
				i = max
			}
			dest = l.find(i)
		}
		l.current = dest
		l.add(a)
		l.add(b)
		l.add(c)
		l.current = current.next
	}

	l.current = l.find(1)
	var sb strings.Builder
	for i := 0; i < l.size-1; i++ {
		sb.WriteString(strconv.Itoa(l.next().label))
	}
	return sb.String()
}

// This is the most lazy solution I could think of. Instead of using find() I'm
// using a slice of labels and their nods in addition to the list so that
// `indices[1]` is the same as `l.find(1)`.
func part2(l *list, moves, cups int) int {
	indices := make([]*node, cups+1)
	for i := 1; i <= l.size; i++ {
		indices[i] = l.find(i)
	}

	current := l.current
	l.current = l.end()
	for i := l.size + 1; i < len(indices); i++ {
		l.add(&node{label: i})
		indices[i] = l.current
	}
	l.current = current

	// We know min is 1 and max is 1_000_000 so min() and max() are is not very
	// useful.
	min, max := l.min().label, l.max().label
	for i := 0; i < moves; i++ {
		current := l.current
		l.next()
		a, b, c := l.take(), l.take(), l.take()
		var dest *node
		for i := current.label - 1; dest == nil; i-- {
			if i < min {
				i = max
			}
			if i == a.label || i == b.label || i == c.label {
				continue
			}
			dest = indices[i]
		}
		l.current = dest
		l.add(a)
		l.add(b)
		l.add(c)
		l.current = current.next
	}

	l.current = indices[1]
	return l.next().label * l.next().label
}
