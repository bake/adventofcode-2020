package main

import "testing"

func testGrid() grid {
	return newGrid([]byte(
		"..##.......\n" +
			"#...#...#..\n" +
			".#....#..#.\n" +
			"..#.#...#.#\n" +
			".#...##..#.\n" +
			"..#.##.....\n" +
			".#.#.#....#\n" +
			".#........#\n" +
			"#.##...#...\n" +
			"#...##....#\n" +
			".#..#...#.#",
	))
}

func TestTraverse(t *testing.T) {
	g := testGrid()
	tt := []struct {
		dx, dy int
		trees  int
	}{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}
	for i, tc := range tt {
		trees := g.traverse(tc.dx, tc.dy)
		if trees != tc.trees {
			t.Fatalf("expected %d. slope to hit %d trees, hit %d", i+1, tc.trees, trees)
		}
	}
}

func TestPart1(t *testing.T) {
	g := testGrid()
	trees := part1(g)
	if trees != 7 {
		t.Fatalf("expected part 1 to return %d trees, got %d", 7, trees)
	}
}

func TestPart2(t *testing.T) {
	g := testGrid()
	trees := part2(g)
	if trees != 336 {
		t.Fatalf("expected part 2 to return %d, got %d", 336, trees)
	}
}
