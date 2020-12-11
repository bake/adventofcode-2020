package main

import (
	"bytes"
	"strings"
)

const (
	cellFloor    = '.'
	cellEmpty    = 'L'
	cellOccupied = '#'
)

type grid struct {
	data          []byte
	width, height int
}

func newGrid(data []byte) *grid {
	var g grid
	data = bytes.TrimSpace(data)
	g.data = bytes.ReplaceAll(data, []byte{'\n'}, []byte{})
	g.width = bytes.IndexByte(data, '\n')
	g.height = len(g.data) / g.width
	return &g
}

func (g *grid) clone() *grid {
	q := grid{
		data:   make([]byte, g.width*g.height),
		width:  g.width,
		height: g.height,
	}
	copy(q.data, g.data)
	return &q
}

func (g *grid) String() string {
	var b strings.Builder
	for y := 0; y < g.height; y++ {
		b.Write(g.data[y*g.width : y*g.width+g.width])
		b.WriteByte('\n')
	}
	return b.String()
}

type evolveFn func(cell byte, adj []byte) (byte, bool)

const (
	distShortsighted = 1
	distInfinity     = -1
)

func (g *grid) evolve(fn evolveFn, dist int) bool {
	next := make([]byte, len(g.data))
	copy(next, g.data)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			i := y*g.width + x
			c := g.data[i]
			c, ok := fn(g.data[i], g.adjacent(x, y, dist))
			if !ok {
				continue
			}
			next[i] = c
		}
	}
	eq := bytes.Compare(g.data, next) == 0
	g.data = next
	return !eq
}

// adjacent returns a slice of cells in all eight directions. dist can be used
// to limit the distance one can see. Use -1 for infinity.
func (g *grid) adjacent(x, y, dist int) []byte {
	var adj []byte
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if x+dx == x && y+dy == y {
				continue
			}
			cell, ok := g.look(x, y, dx, dy, dist)
			if !ok {
				continue
			}
			adj = append(adj, cell)
		}
	}
	return adj
}

// look in a direction and return the first non floor cell.
func (g *grid) look(x, y, dx, dy, dist int) (cell byte, ok bool) {
	for ; dist != 0; dist-- {
		x, y = x+dx, y+dy
		if x < 0 || g.width <= x {
			break
		}
		if y < 0 || g.height <= y {
			break
		}
		if g.data[y*g.width+x] != cellFloor {
			return g.data[y*g.width+x], true
		}
	}
	return cellEmpty, false
}
