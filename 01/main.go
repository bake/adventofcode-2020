// This is the boring solution as it just uses one for loops per number. No need
// to read any further.
//
// The number of loops could have been reduced by using a map and just checking
// if dest - current number exists but this would have been too much on the
// first day. Similary the calculations done in the first part could have been
// stored and reused in the second but, again, this is the first day.

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	dest := flag.Int("dest", 2020, "The sum you're looking for")
	flag.Parse()

	nums, err := input(os.Stdin)
	if err != nil {
		return err
	}

	out, err := part1(nums, *dest)
	if err != nil {
		return err
	}
	fmt.Println(out)

	out, err = part2(nums, *dest)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func part1(nums []int, dest int) (int, error) {
	for i, m := range nums {
		for _, n := range nums[:i] {
			if m+n != dest {
				continue
			}
			return m * n, nil
		}
	}
	return 0, errors.New("nothing found")
}

func part2(nums []int, dest int) (int, error) {
	for i, m := range nums {
		for j, n := range nums[:i] {
			for _, o := range nums[:j] {
				if m+n+o != dest {
					continue
				}
				return m * n * o, nil
			}
		}
	}
	return 0, errors.New("nothing found")
}

func input(r io.Reader) ([]int, error) {
	var nums []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}
