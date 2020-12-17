package main

type point2d struct{ x, y int }

type bounds2d struct{ min, max point2d }

type rectangle struct {
	bounds bounds2d
	data   map[point2d]interface{}
}

func (r *rectangle) neighbors(p point2d) []point2d {
	var ps []point2d
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			q := point2d{p.x + dx, p.y + dy}
			if _, ok := r.data[q]; ok {
				ps = append(ps, q)
			}
		}
	}
	return ps
}

func (r *rectangle) evolve(fn evolveFn) {
	data := map[point2d]interface{}{}
	r.bounds.min.x--
	r.bounds.min.y--
	r.bounds.max.x++
	r.bounds.max.y++
	for y := r.bounds.min.y; y <= r.bounds.max.y; y++ {
		for x := r.bounds.min.x; x <= r.bounds.max.x; x++ {
			p := point2d{x, y}
			_, active := r.data[p]
			if fn(active, len(r.neighbors(p))) {
				data[p] = nil
			}
		}
	}
	r.data = data
}
