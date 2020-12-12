package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
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
	as, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(as))
	fmt.Println(part2(as))
	return nil
}

func input(r io.Reader) ([]action, error) {
	var as []action
	s := bufio.NewScanner(r)
	for s.Scan() {
		var a action
		if _, err := fmt.Sscanf(s.Text(), "%c%d", &a.dir, &a.steps); err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, s.Err()
}

func part1(as []action) int {
	f := newFerry(image.Pt(0, 0), image.Pt(1, 0))
	for _, a := range as {
		switch a.dir {
		case left:
			f.dir = f.rotate(f.dir, a.steps*3)
		case right:
			f.dir = f.rotate(f.dir, a.steps)
		case forward:
			f.pos = f.pos.Add(f.dir.Mul(a.steps))
		}
		f.move(a)
	}
	return int(math.Abs(float64(f.pos.X))) + int(math.Abs(float64(f.pos.Y)))
}

func part2(as []action) int {
	f := newFerry(image.Pt(0, 0), image.Pt(1, 0))
	p := newFerry(image.Pt(10, -1), image.Pt(0, 0))
	for _, a := range as {
		p.move(a)
		switch a.dir {
		case left:
			for i := 0; i < a.steps; i += 90 {
				p.pos.X, p.pos.Y = p.pos.Y, -p.pos.X
			}
		case right:
			for i := 0; i < a.steps; i += 90 {
				p.pos.X, p.pos.Y = -p.pos.Y, p.pos.X
			}
		case forward:
			f.pos = f.pos.Add(p.pos.Mul(a.steps))
		}
	}
	return int(math.Abs(float64(f.pos.X))) + int(math.Abs(float64(f.pos.Y)))
}

type direction byte

const (
	east    direction = 'E'
	south             = 'S'
	west              = 'W'
	north             = 'N'
	left              = 'L'
	right             = 'R'
	forward           = 'F'
)

type action struct {
	dir   direction
	steps int
}

func (a action) String() string { return fmt.Sprintf("%c%d", a.dir, a.steps) }

type ferry struct {
	pos image.Point
	dir image.Point
}

func newFerry(pos, dir image.Point) *ferry { return &ferry{pos, dir} }

func (f *ferry) String() string {
	return fmt.Sprintf("(%d,%d)->(%d,%d)", f.pos.X, f.pos.Y, f.dir.X, f.dir.Y)
}

func (f *ferry) move(a action) {
	switch a.dir {
	case east:
		f.pos.X += a.steps
	case south:
		f.pos.Y += a.steps
	case west:
		f.pos.X -= a.steps
	case north:
		f.pos.Y -= a.steps
	}
}

func (f *ferry) rotate(p image.Point, deg int) image.Point {
	for i := 0; i < deg; i += 90 {
		p.X, p.Y = -p.Y, p.X
	}
	return p
}
