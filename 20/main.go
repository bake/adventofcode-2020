package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ts, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(ts))
	return nil
}

func input(r io.Reader) (map[int]tile, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	ts := map[int]tile{}
	for _, t := range bytes.Split(data, []byte{'\n', '\n'}) {
		parts := bytes.SplitN(t, []byte{'\n'}, 2)
		var i int
		if _, err := fmt.Sscanf(string(parts[0]), "Tile %d:", &i); err != nil {
			return nil, err
		}
		ts[i] = bytes.Split(bytes.TrimSpace(parts[1]), []byte{'\n'})
	}
	return ts, nil
}

type tile [][]byte

// rotate a tile 90deg clockwise.
func (t tile) rotate() tile {
	w, h := len(t), len(t[0])
	dst := make(tile, w)
	for x := range dst {
		dst[x] = make([]byte, h)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dst[y][w-1-x] = t[x][y]
		}
	}
	return dst
}

func (t tile) flipX() tile {
	w := len(t)
	dst := make(tile, w)
	for y := range t {
		dst[y] = make([]byte, w)
		for i, j := 0, w-1; i < j; i, j = i+1, j-1 {
			dst[y][i], dst[y][j] = t[y][j], t[y][i]
		}
	}
	return dst
}

func (t tile) flipY() tile {
	h := len(t[0])
	dst := make(tile, h)
	for i, j := 0, h-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = t[j], t[i]
	}
	return dst
}

func (t tile) edge() []byte { return t[0] }

func (t tile) String() string { return string(bytes.Join(t, []byte{'\n'})) }

// list is a helper structure for duplicate free slices.
type list map[int]interface{}

func (l list) add(i int) { l[i] = nil }

func (l list) slice() []int {
	var is []int
	for i := range l {
		is = append(is, i)
	}
	return is
}

type edge struct {
	tile         int
	rotations    int
	flipX, flipY bool
}

func edges(ts map[int]tile) (map[string][]edge, map[int]list) {
	es := map[string][]edge{}
	ns := map[int]list{}
	for i, t := range ts {
		ns[i] = list{}
		for r := 0; r < 4; r++ {
			e := string(t.edge())
			es[e] = append(es[e], edge{tile: i, rotations: r})

			e = string(t.flipX().edge())
			es[e] = append(es[e], edge{tile: i, rotations: r, flipX: true})
			e = string(t.flipY().edge())
			es[e] = append(es[e], edge{tile: i, rotations: r, flipY: true})

			for _, v := range es[e] {
				ns[i].add(v.tile)
				ns[v.tile].add(i)
			}

			t = t.rotate()
		}
	}
	return es, ns
}

func part1(ts map[int]tile) int {
	_, ns := edges(ts)
	prod := 1
	for i, l := range ns {
		if len(l) == 3 {
			prod *= i
		}
	}
	return prod
}
