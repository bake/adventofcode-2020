package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	nums, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(nums, 2020))
	fmt.Println(part1(nums, 30000000))
	return nil
}

func input(r io.Reader) ([]int, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var nums []int
	for _, n := range strings.Split(string(raw), ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// As you can see, there is no part 2. I'm lazy and can live with ~3s.
func part1(nums []int, dst int) int {
	dists := map[int][]int{}
	for i, n := range nums {
		dists[n] = append(dists[n], i+1)
	}
	var last int
	for i := len(nums) + 1; i <= dst; i++ {
		dist := dists[last]
		if len(dist) < 2 {
			last, dists[0] = 0, append(dists[0], i)
			continue
		}
		dists[last] = dist[len(dist)-2:]
		num := dist[len(dist)-1] - dist[len(dist)-2]
		last, dists[num] = num, append(dists[num], i)
	}
	return last
}
