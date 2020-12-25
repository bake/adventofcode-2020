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
	rs, t, ts, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(rs, ts))
	fmt.Println(part2(rs, t, ts))
	return nil
}

func part1(rs map[string]rules, ts []ticket) int {
	var sum int
	for _, t := range ts {
		for _, v := range t {
			if !validate(rs, v) {
				sum += v
			}
		}
	}
	return sum
}

func part2(rsm map[string]rules, t ticket, ts []ticket) int {
	// Remove invalid tickets.
	var vts []ticket
	for _, t := range ts {
		valid := true
		for _, v := range t {
			if !validate(rsm, v) {
				valid = false
				break
			}
		}
		if valid {
			vts = append(vts, t)
		}
	}
	ts = vts

	// fields is a map from rule name to a list of possible columns.
	fields := map[string]ticket{}
	for n, rs := range rsm {
		for i := range ts[0] {
			valid := true
			for _, t := range ts {
				if !rs.valid(t[i]) {
					valid = false
					break
				}
			}
			if valid {
				fields[n] = append(fields[n], i)
			}
		}
	}

	// Match fields.
	res := map[string]int{}
	for len(res) < len(rsm) {
		var found int
		for n, rs := range fields {
			if len(rs) == 1 {
				found, res[n] = rs[0], rs[0]
				break
			}
		}
		for n, t := range fields {
			fields[n] = t.delete(found)
			if len(fields[n]) == 0 {
				delete(fields, n)
			}
		}
	}

	prod := 1
	for n, i := range res {
		if strings.HasPrefix(n, "departure") {
			prod *= t[i]
		}
	}
	return prod
}

func validate(rs map[string]rules, v int) bool {
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

type rules []rule

func (rs rules) valid(v int) bool {
	for _, r := range rs {
		if r.valid(v) {
			return true
		}
	}
	return false
}

type ticket []int

func (t ticket) delete(v int) ticket {
	for i, w := range t {
		if v != w {
			continue
		}
		return append(t[:i], t[i+1:]...)
	}
	return t
}

func input(r io.Reader) (rsm map[string]rules, t ticket, ts []ticket, err error) {
	rsm = map[string]rules{}
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
			rsm[n] = append(rsm[n], rs...)
		case "your ticket":
			t, err = inputInts(s.Text())
			if err != nil {
				return nil, t, ts, err
			}
		case "nearby tickets":
			nt, err := inputInts(s.Text())
			if err != nil {
				return nil, t, ts, err
			}
			ts = append(ts, nt)
		}
	}
	return rsm, t, ts, s.Err()
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
