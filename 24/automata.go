package main

type automata struct {
	grid *grid
}

func newAutomata(g *grid) *automata {
	return &automata{grid: g}
}

func (a *automata) step() {
	cpy := a.grid.copy()
	for p := range a.grid.cells {
		for _, q := range a.grid.neighbours(p) {
			if a.grid.at(q) == black {
				continue
			}
			a.grid.set(q, white)
		}
	}
	for p := range a.grid.cells {
		c := a.grid.at(p)
		b, _ := a.grid.neighbourColors(p)
		if c == black && (b == 0 || b > 2) {
			cpy.set(p, white)
		}
		if c == white && b == 2 {
			cpy.set(p, black)
		}
	}
	a.grid.cells = cpy.cells
}
