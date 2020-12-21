package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	is, as, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(is, as))
	fmt.Println(part2(is, as))
	return nil
}

type combination map[string][]list

func input(r io.Reader) (combination, combination, error) {
	is := combination{}
	as := combination{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		sides := strings.SplitN(s.Text(), "contains", 2)
		left := strings.Trim(sides[0], "() ")
		right := strings.Trim(sides[1], "() ")
		for _, i := range strings.Fields(left) {
			is[i] = append(is[i], strings.Split(right, ", "))
		}
		for _, a := range strings.Split(right, ", ") {
			as[a] = append(as[a], strings.Fields(left))
		}
	}
	return is, as, s.Err()
}

func part1(is combination, as combination) int {
	var bad list
	for _, ls := range as {
		found := ls[0]
		for _, l := range ls {
			found = found.intersect(l)
		}
		bad = append(bad, found...)
	}
	bad = bad.unique()

	var sum int
	for i := range is {
		if bad.contains(i) {
			continue
		}
		sum += len(is[i])
	}
	return sum
}

type ingredient struct{ name, allergene string }

type byAllergene []ingredient

func (is byAllergene) Len() int           { return len(is) }
func (is byAllergene) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is byAllergene) Less(i, j int) bool { return is[i].allergene < is[j].allergene }

func part2(is combination, as combination) string {
	unique := map[string]list{}
	for a, ls := range as {
		found := ls[0]
		for _, l := range ls {
			found = found.intersect(l)
		}
		unique[a] = append(unique[a], found...)
	}

	var res []ingredient

	for {
		var sum int
		for _, l := range unique {
			if len(l) > 0 {
				sum++
			}
		}
		if sum == 0 {
			break
		}

		var found string
		for a, l := range unique {
			if len(l) == 1 {
				found = l[0]
				res = append(res, ingredient{found, a})
				break
			}
		}

		for a, l := range unique {
			unique[a] = l.delete(found)
		}
	}

	sort.Sort(byAllergene(res))
	var names []string
	for _, i := range res {
		names = append(names, i.name)
	}
	return strings.Join(names, ",")
}
