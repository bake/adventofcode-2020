package main

import "math"

type point struct{ x, y int }

func (p point) add(q point) point {
	p.x += q.x
	p.y += q.y
	return p
}

type color bool

const (
	black color = true
	white color = false
)

type grid struct {
	min, max point
	cells    map[point]color
}

func newGrid() *grid {
	return &grid{
		min:   point{math.MaxInt64, math.MaxInt64},
		max:   point{math.MinInt64, math.MinInt64},
		cells: map[point]color{},
	}
}

func (g *grid) copy() *grid {
	q := newGrid()
	q.min = g.min
	q.max = g.max
	for k, v := range g.cells {
		q.cells[k] = v
	}
	return q
}

func (g *grid) at(p point) color {
	return g.cells[p]
}

func (g *grid) set(p point, c color) {
	g.cells[p] = c
}

// flip a cell in the hexagonal grid. As many others, I copied the coordinates
// from https://www.redblobgames.com/grids/hexagons/#neighbors-doubled.
func (g *grid) flip(path []direction) {
	var p point
	for _, d := range path {
		switch d {
		case "e":
			p.x += 2
		case "se":
			p.x++
			p.y++
		case "sw":
			p.x--
			p.y++
		case "w":
			p.x -= 2
		case "nw":
			p.x--
			p.y--
		case "ne":
			p.x++
			p.y--
		}
	}
	if p.x < g.min.x {
		g.min.x = p.x
	}
	if p.x > g.max.x {
		g.max.x = p.x
	}
	if p.y < g.min.y {
		g.min.y = p.y
	}
	if p.y > g.max.y {
		g.max.y = p.y
	}
	g.cells[p] = !g.cells[p]
}

func (g *grid) neighbours(p point) []point {
	ds := []point{
		{2, 0},
		{1, 1},
		{-1, 1},
		{-2, 0},
		{-1, -1},
		{1, -1},
	}
	var ns []point
	for _, d := range ds {
		ns = append(ns, p.add(d))
	}
	return ns
}

func (g *grid) neighbourColors(p point) (b int, w int) {
	cs := map[color]int{}
	for _, q := range g.neighbours(p) {
		cs[g.at(q)]++
	}
	return cs[black], cs[white]
}
