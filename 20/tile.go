package main

import (
	"bytes"
)

const (
	north = iota
	east
	south
	west
)

type tile [][]byte

func newTile(width, height int) tile {
	t := make([][]byte, height)
	for y := range t {
		t[y] = make([]byte, width)
	}
	return t
}

func (t tile) rotate() tile {
	w, h := len(t), len(t[0])
	dst := newTile(w, h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dst[y][w-1-x] = t[x][y]
		}
	}
	return dst
}

func (t tile) flipY() tile {
	w, h := len(t), len(t[0])
	dst := newTile(w, h)
	for i, j := 0, h-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = t[j], t[i]
	}
	return dst
}

// trim removes the outermost pixels of a tile.
func (t tile) trim() tile {
	w, h := len(t), len(t[0])
	dst := newTile(w-2, h-2)
	for y := range dst {
		dst[y] = t[y+1][1 : w-1]
	}
	return dst
}

func (t tile) edge(side int) (string, bool) {
	switch side {
	case north:
		return string(t[0]), true
	case south:
		return string(t[len(t)-1]), true
	case west:
		var edge string
		for _, row := range t {
			edge += string(row[0])
		}
		return edge, true
	case east:
		var edge string
		for _, row := range t {
			edge += string(row[len(row)-1])
		}
		return edge, true
	}
	return "", false
}

// orientations returns a tile in all possible rotations and flips.
func (t tile) orientations() []tile {
	ts := []tile{t, t.flipY()}
	for r := 0; r < 4; r++ {
		t = t.rotate()
		ts = append(ts, t, t.flipY())
	}
	return ts
}

func (t tile) String() string { return string(bytes.Join(t, []byte{'\n'})) }
