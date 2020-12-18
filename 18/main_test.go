package main

import "testing"

func TestEval(t *testing.T) {
	tt := []struct {
		eq  string
		val int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, tc := range tt {
		_, val := eval(tc.eq)
		if val != tc.val {
			t.Fatalf("expected %q to evaluate to %d, got %d", tc.eq, tc.val, val)
		}
	}
}

func TestEval2(t *testing.T) {
	tt := []struct {
		eq  string
		val int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}
	for _, tc := range tt {
		val := eval2(tc.eq)
		if val != tc.val {
			t.Fatalf("expected %q to evaluate to %d, got %d", tc.eq, tc.val, val)
		}
	}
}
