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
	rs, _, ts, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(rs, ts))
	return nil
}

func part1(rs map[string][]rule, ts []ticket) int {
	var sum int
	for _, t := range ts {
		for _, v := range t {
			if !valid(rs, v) {
				sum += v
			}
		}
	}
	return sum
}

func valid(rs map[string][]rule, v int) bool {
	for _, rs := range rs {
		for _, r := range rs {
			if r.valid(v) {
				return true
			}
		}
	}
	return false
}

type rule struct{ min, max int }

func (r rule) valid(v int) bool {
	return r.min <= v && v <= r.max
}

type ticket []int

func input(r io.Reader) (rules map[string][]rule, t ticket, ts []ticket, err error) {
	rules = map[string][]rule{}
	s := bufio.NewScanner(r)
	step := "rules"
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		if strings.HasSuffix(s.Text(), ":") {
			step = strings.TrimSuffix(s.Text(), ":")
			continue
		}
		switch step {
		case "rules":
			n, rs, err := inputRules(s.Text())
			if err != nil {
				return nil, t, ts, err
			}
			rules[n] = append(rules[n], rs...)
		case "your ticket":
			t, err = inputInts(s.Text())
			if err != nil {
				return nil, t, ts, err
			}
		case "nearby tickets":
			t, err = inputInts(s.Text())
			if err != nil {
				return nil, t, ts, err
			}
			ts = append(ts, t)
		}
	}
	return rules, t, ts, s.Err()
}

func inputInts(text string) ([]int, error) {
	var is []int
	for _, raw := range strings.Split(text, ",") {
		i, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		is = append(is, i)
	}
	return is, nil
}

func inputRules(text string) (string, []rule, error) {
	rp := NewRuleParser(text)
	for rp.Parse() {
	}
	return rp.Name(), rp.Rules(), rp.Err()
}
