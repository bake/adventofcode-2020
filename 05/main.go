package main

import (
	"bufio"
	"fmt"
	"io"
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
	passes, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(passes))
	fmt.Println(part2(passes))
	return nil
}

func input(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	var lines []string
	for s.Scan() {
		if s.Text() != "" {
			lines = append(lines, s.Text())
		}
	}
	return lines, s.Err()
}

func part1(passes []string) int {
	var max int
	for _, pass := range passes {
		if s := seat(pass); s > max {
			max = s
		}
	}
	return max
}

func part2(passes []string) int {
	taken := make([]bool, 1024)
	for _, pass := range passes {
		taken[seat(pass)] = true
	}
	for i := 1; i < len(taken)-1; i++ {
		if taken[i-1] && !taken[i] && taken[i+1] {
			return i
		}
	}
	return 0
}

func splitPass(pass string) (row, column string) {
	i := strings.LastIndexAny(pass, "FB")
	if i < 0 {
		return "", pass
	}
	// This line drove me crazy. I initially wrote `pass[:i], pass[i+1:]` which
	// (now) obviously removed the rows last bit.
	return pass[:i+1], pass[i+1:]
}

func seat(pass string) int {
	row, column := splitPass(pass)
	a, _ := search(0, 127, row)
	b, _ := search(0, 7, column)
	return a*8 + b
}

func search(min, max int, queue string) (int, error) {
	if min >= max {
		return min, nil
	}
	if len(queue) == 0 {
		return 0, fmt.Errorf("unexpected end of queue")
	}
	med := (min + max) / 2
	switch queue[0] {
	case 'F', 'L':
		return search(min, med, queue[1:])
	case 'B', 'R':
		return search(med+1, max, queue[1:])
	default:
		return 0, fmt.Errorf("unexpected direction %q", queue[0])
	}
}
