package main

import "testing"

func TestPart1(t *testing.T) {
	tt := []struct {
		nums []int
		dest int
		out  int
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 514579},
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
		{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 241861950},
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
