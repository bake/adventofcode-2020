package main

import (
	"bufio"
	"fmt"
	"io"
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
	p1, p2, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(p1, p2))
	fmt.Println(part2(p1, p2))
	return nil
}

type player []int

func (p player) score() int {
	var sum int
	for i, v := range p {
		sum += (len(p) - i) * v
	}
	return sum
}

func (p player) encode() string {
	var sb strings.Builder
	for _, i := range p {
		sb.WriteString(strconv.Itoa(i) + ",")
	}
	return sb.String()
}

func (p player) copy() player {
	q := make(player, len(p))
	copy(q, p)
	return q
}

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
	if len(p1) > len(p2) {
		return p1.score()
	}
	return p2.score()
}

func recursiveCombat(p1, p2 player) (bool, int) {
	seen := map[string]interface{}{}
	for len(p1) > 0 && len(p2) > 0 {
		key := p1.encode() + "|" + p2.encode()
		if _, ok := seen[key]; ok {
			return true, p1.score()
		}
		seen[key] = nil

		winner := p1[0] > p2[0]
		if p1[0] < len(p1) && p2[0] < len(p2) {
			winner, _ = recursiveCombat(p1.copy()[1:p1[0]+1], p2.copy()[1:p2[0]+1])
		}
		switch winner {
		case true:
			p1 = append(p1[1:], p1[0])
			p1, p2 = append(p1, p2[0]), p2[1:]
		case false:
			p2 = append(p2[1:], p2[0])
			p1, p2 = p1[1:], append(p2, p1[0])
		}
	}

	if len(p1) > len(p2) {
		return true, p1.score()
	}
	return false, p2.score()
}

func part2(p1, p2 player) int {
	_, score := recursiveCombat(p1, p2)
	return score
}
