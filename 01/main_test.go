package main

import "testing"

func TestPart1(t *testing.T) {
	tt := []struct {
		nums []int
		dest int
		out  int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 24},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 10, 24},
		{[]int{1, 1, 1, 1, 1, 9, 9, 9, 9, 9}, 10, 9},
	}
	for i, tc := range tt {
		out, err := part1(tc.nums, tc.dest)
		if err != nil {
			t.Fatal(err)
		}
		if out != tc.out {
			t.Fatalf("expected %d. case to return %d, got %d", i+1, tc.out, out)
		}
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		nums []int
		dest int
		out  int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 30},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 10, 30},
	}
	for i, tc := range tt {
		out, err := part2(tc.nums, tc.dest)
		if err != nil {
			t.Fatal(err)
		}
		if out != tc.out {
			t.Fatalf("expected %d. case to return %d, got %d", i+1, tc.out, out)
		}
	}
}
