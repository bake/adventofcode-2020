package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type password struct {
	a, b   int
	letter rune
	pw     string
}

func (p password) String() string { return p.pw }

func (p password) validateLength() bool {
	var num int
	for _, r := range p.pw {
		if r != p.letter {
			continue
		}
		num++
	}
	return p.a <= num && num <= p.b
}

func (p password) validatePosition() bool {
	rs := []rune(p.pw)
	a, b := rs[p.a-1], rs[p.b-1]
	if a == p.letter && b != p.letter {
		return true
	}
	if a != p.letter && b == p.letter {
		return true
	}
	return false
}

func run() error {
	pws, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(pws))
	fmt.Println(part2(pws))
	return nil
}

func input(r io.Reader) ([]password, error) {
	var pws []password
	s := bufio.NewScanner(r)
	for s.Scan() {
		var pw password
		_, err := fmt.Sscanf(s.Text(), "%d-%d %c: %s", &pw.a, &pw.b, &pw.letter, &pw.pw)
		if err != nil {
			return nil, err
		}
		pws = append(pws, pw)
	}
	return pws, nil
}

func part1(pws []password) int {
	var num int
	for _, pw := range pws {
		if !pw.validateLength() {
			continue
		}
		num++
	}
	return num
}

func part2(pws []password) int {
	var num int
	for _, pw := range pws {
		if !pw.validatePosition() {
			continue
		}
		num++
	}
	return num
}
