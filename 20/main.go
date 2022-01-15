package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
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
	fmt.Println(part2(ts))
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

// edges returns an map from tile ids to their connected tile ids.
func edges(ts map[int]tile) map[int][]int {
	es := map[string][]int{}
	ns := map[int]map[int]interface{}{}
	for i, t := range ts {
		ns[i] = map[int]interface{}{}
		for r := 0; r < 4; r++ {
			e, ok := t.edge(north)
			if !ok {
				continue
			}
			es[e] = append(es[e], i)
			if e, ok := t.flipY().edge(north); ok {
				es[e] = append(es[e], i)
			}
			for _, v := range es[e] {
				if i == v {
					continue
				}
				ns[i][v] = nil
				ns[v][i] = nil
			}
			t = t.rotate()
		}
	}

	res := map[int][]int{}
	for id, ids := range ns {
		for i := range ids {
			res[id] = append(res[id], i)
		}
	}
	return res
}

func part1(ts map[int]tile) int {
	prod := 1
	for i, l := range edges(ts) {
		if len(l) == 2 {
			prod *= i
		}
	}
	return prod
}

// neighbour returns an adjacent tile.
func neighbour(ts map[int]tile, stop map[int]interface{}, id int, side1, side2 int) (int, tile, bool) {
	e1, ok := ts[id].edge(side1)
	if !ok {
		return 0, nil, false
	}
	for i, t := range ts {
		if i == id {
			continue
		}
		if _, ok := stop[i]; ok {
			continue
		}
		for _, ta := range t.orientations() {
			if e2, ok := ta.edge(side2); ok && e1 == e2 {
				ts[i] = ta
				return i, ta, true
			}
		}
	}
	return 0, nil, false
}

func tiledRow(ts map[int]tile, used map[int]interface{}, id, width int) ([]int, []tile, bool) {
	ids := []int{id}
	tiles := []tile{ts[id]}
	for x := len(ids); x < width; x++ {
		var ok bool
		id, _, ok = neighbour(ts, used, id, east, west)
		if !ok {
			return ids, tiles, false
		}
		ids = append(ids, id)
		used[id] = nil
	}
	return ids, tiles, true
}

func tiledImage(ts map[int]tile, start int) ([][]int, bool) {
	size := int(math.Sqrt(float64(len(ts))))
	var ids [][]int
	used := map[int]interface{}{}
	for len(ts) > 0 {
		row, _, ok := tiledRow(ts, used, start, size)
		if !ok {
			break
		}
		ids = append(ids, row)
		start, _, _ = neighbour(ts, used, start, south, north)
	}
	if len(ids) == size && len(ids[size-1]) == size {
		return ids, true
	}
	return nil, false
}

// patterns returns the number of times a given pattern occures in a tile.
func patterns(t, pattern tile) int {
	var num int
	for yt := 0; yt < len(t)-len(pattern); yt++ {
		for xt := 0; xt < len(t[0])-len(pattern[0]); xt++ {
			found := true
			for ym := range pattern {
				for xm := range pattern[ym] {
					if pattern[ym][xm] == ' ' {
						continue
					}
					if pattern[ym][xm] != t[yt+ym][xt+xm] {
						found = false
						break
					}
				}
			}
			if found {
				num++
			}
		}
	}
	return num
}

func part2(ts map[int]tile) int {
	var ids [][]int
	var ok bool
	for i, l := range edges(ts) {
		if len(l) != 2 {
			continue
		}
		ids, ok = tiledImage(ts, i)
		if ok {
			break
		}
	}

	image := newTile(0, len(ids)*8)
	for row := range ids {
		for col := range ids[row] {
			id := ids[row][col]
			t := ts[id].trim()
			for y := 0; y < len(t); y++ {
				yi := row*len(t) + y
				image[yi] = append(image[yi], t[y]...)
			}
		}
	}

	var seaHashes int
	for _, row := range image {
		for _, cell := range row {
			if cell == '#' {
				seaHashes++
			}
		}
	}

	monster := tile{
		[]byte("                  # "),
		[]byte("#    ##    ##    ###"),
		[]byte(" #  #  #  #  #  #   "),
	}
	var monsterHashes int
	for _, row := range monster {
		for _, cell := range row {
			if cell == '#' {
				monsterHashes++
			}
		}
	}
	var monsters int
	for _, t := range image.orientations() {
		num := patterns(t, monster)
		if num > monsters {
			monsters = num
		}
	}

	return seaHashes - monsters*monsterHashes
}
