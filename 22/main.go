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
	p1, p2, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(p1, p2))
	return nil
}

type player []int

func input(r io.Reader) (player, player, error) {
	var ps []player
	s := bufio.NewScanner(r)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		if strings.HasPrefix(s.Text(), "Player") {
			ps = append(ps, player{})
			continue
		}
		var n int
		if _, err := fmt.Sscanf(s.Text(), "%d", &n); err != nil {
			return player{}, player{}, err
		}
		ps[len(ps)-1] = append(ps[len(ps)-1], n)
	}
	return ps[0], ps[1], nil
}

func part1(p1, p2 player) int {
	for len(p1) > 0 && len(p2) > 0 {
		switch p1[0] > p2[0] {
		case true:
			p1 = append(p1[1:], p1[0])
			p1, p2 = append(p1, p2[0]), p2[1:]
		case false:
			p2 = append(p2[1:], p2[0])
			p1, p2 = p1[1:], append(p2, p1[0])
		}
	}
	var winner []int
	switch len(p1) > len(p2) {
	case true:
		winner = p1
	case false:
		winner = p2
	}
	var sum int
	for i := 0; i < len(winner); i++ {
		sum += (len(winner) - i) * winner[i]
	}
	return sum
}
