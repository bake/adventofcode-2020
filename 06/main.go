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
	answers, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(answers))
	fmt.Println(part2(answers))
	return nil
}

func input(r io.Reader) ([]string, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(raw), "\n\n"), nil
}

func part1(answers []string) int {
	var sum int
	for _, a := range answers {
		letters := make([]bool, 'z'-'a'+1)
		for _, r := range a {
			if 'a' <= r && r <= 'z' {
				letters[r-'a'] = true
			}
		}
		for _, ok := range letters {
			if ok {
				sum++
			}
		}
	}
	return sum
}

func part2(answers []string) int {
	var sum int
	for _, answer := range answers {
		lines := strings.Count(strings.Trim(answer, "\n"), "\n") + 1
		for _, r := range strings.SplitN(answer, "\n", 2)[0] {
			if strings.Count(answer, string(r)) == lines {
				sum++
			}
		}
	}
	return sum
}
