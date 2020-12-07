package main

import (
	"bufio"
	"flag"
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
	color := flag.String("color", "shiny gold", "The bag color")
	flag.Parse()

	bags, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(*color, bags))
	fmt.Println(part2(*color, bags))
	return nil
}

type luggage struct {
	num   int
	color string
}

func input(r io.Reader) (map[string][]luggage, error) {
	m := map[string][]luggage{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		color, bags := parseBag(s.Text())
		m[color] = append(m[color], bags...)
	}
	return m, s.Err()
}

func parseBag(line string) (color string, bags []luggage) {
	parts := strings.SplitN(line, "bags contain", 2)
	if len(parts) != 2 {
		return "", nil
	}
	for _, bag := range strings.Split(strings.Trim(parts[1], "."), ",") {
		bag = strings.TrimSpace(bag)
		var num int
		var mod, color string
		_, err := fmt.Sscanf(bag, "%d %s %s bag", &num, &mod, &color)
		if err != nil {
			continue
		}
		bags = append(bags, luggage{num, mod + " " + color})
	}
	return strings.TrimSpace(parts[0]), bags
}

func part1(color string, luggage map[string][]luggage) int {
	// Invert the graph to increase lookup speed.
	colors := map[string][]string{}
	for k, vs := range luggage {
		for _, v := range vs {
			colors[v.color] = append(colors[v.color], k)
		}
	}
	matches := map[string]interface{}{}
	queue := []string{color}
	for len(queue) > 0 {
		color, queue = queue[0], queue[1:]
		if _, ok := matches[color]; ok {
			continue
		}
		matches[color] = nil
		for _, c := range colors[color] {
			queue = append(queue, c)
		}
	}
	return len(matches) - 1
}

func part2(color string, bags map[string][]luggage) int {
	return count(color, bags) - 1
}

func count(color string, lug map[string][]luggage) int {
	num := 1
	for _, bag := range lug[color] {
		num += bag.num * count(bag.color, lug)
	}
	return num
}
