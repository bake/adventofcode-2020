package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	min, times, err := input(os.Stdin)
	if err != nil {
		return err
	}
	cpy := make([]int, len(times))
	copy(cpy, times)
	fmt.Println(part1(min, cpy))
	copy(cpy, times)
	fmt.Println(part2(cpy))
	return nil
}

func input(r io.Reader) (int, []int, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, nil, err
	}
	lines := bytes.Split(data, []byte{'\n'})
	min, err := strconv.Atoi(string(lines[0]))
	if err != nil {
		return 0, nil, err
	}
	var times []int
	for _, t := range bytes.Split(lines[1], []byte{','}) {
		// I'd usually ignore "x"s in the input but part 2 needs to know about them.
		time, err := strconv.Atoi(string(t))
		if err != nil {
			time = 0
		}
		times = append(times, time)
	}
	return min, times, nil
}

func part1(arrival int, times []int) int {
	sort.Ints(times)
	min, time := math.MaxInt64, 0
	for _, t := range times {
		if t == 0 {
			continue
		}
		diff := t * int(math.Ceil(float64(arrival)/float64(t)))
		if diff%arrival < min {
			min, time = diff%arrival, t
		}
	}
	return min * time
}

func part2(times []int) int {
	min, prod := 0, 1
	for i, bus := range times {
		if bus == 0 {
			continue
		}
		for (min+i)%bus != 0 {
			min += prod
		}
		prod *= bus
	}
	return min
}
