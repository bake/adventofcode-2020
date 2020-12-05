package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tt := []struct {
		pass     string
		min, max int
		value    int
	}{
		{"BFFFBBF", 0, 127, 70},
		{"RRR", 0, 7, 7},
		{"FFFBBBF", 0, 127, 14},
		{"BBFFBBF", 0, 127, 102},
		{"RLL", 0, 7, 4},
		{"FFFFBBF", 0, 127, 6},
		{"FFFFBBB", 0, 127, 7},
		{"LLR", 0, 7, 1},
	}
	for _, tc := range tt {
		value, err := search(tc.min, tc.max, tc.pass)
		if err != nil {
			t.Fatal(err)
		}
		if value != tc.value {
			t.Fatalf("expected pass %s to have a value of %d, got %d", tc.pass, tc.value, value)
		}
	}
}

func TestSeat(t *testing.T) {
	tt := []struct {
		pass string
		seat int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"FFFFBBFLLR", 49},
		{"FFFFBBBLLR", 57},
	}
	for _, tc := range tt {
		s := seat(tc.pass)
		if s != tc.seat {
			t.Fatalf("expected pass %s to have seat %d, got %d", tc.pass, tc.seat, s)
		}
	}
}
