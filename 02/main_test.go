package main

import "testing"

func TestValidateLength(t *testing.T) {
	tt := []struct {
		pw    password
		valid bool
	}{
		{pw: password{1, 3, 'a', "abcde"}, valid: true},
		{pw: password{1, 3, 'b', "cdefg"}, valid: false},
		{pw: password{2, 9, 'c', "ccccccccc"}, valid: true},
	}
	for _, tc := range tt {
		if tc.pw.validateLength() != tc.valid {
			t.Fatalf("expected %s to be %v, got %v", tc.pw, tc.valid, !tc.valid)
		}
	}
}

func TestValidatePosition(t *testing.T) {
	tt := []struct {
		pw    password
		valid bool
	}{
		{pw: password{1, 3, 'a', "abcde"}, valid: true},
		{pw: password{1, 3, 'b', "cdefg"}, valid: false},
		{pw: password{2, 9, 'c', "ccccccccc"}, valid: false},
	}
	for _, tc := range tt {
		if tc.pw.validatePosition() != tc.valid {
			t.Fatalf("expected %s to be %v, got %v", tc.pw, tc.valid, !tc.valid)
		}
	}
}
