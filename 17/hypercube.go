package main

type point4d struct{ x, y, z, w int }

type bounds4d struct{ min, max point4d }

type hypercube struct {
	bounds bounds4d
	data   map[point4d]interface{}
}

func newHypercube(c *cube) *hypercube {
	data := map[point4d]interface{}{}
	for k := range c.data {
		data[point4d{x: k.x, y: k.y, z: k.z, w: 0}] = nil
	}
	return &hypercube{
		bounds: bounds4d{
			min: point4d{x: c.bounds.min.x, y: c.bounds.min.y, z: c.bounds.min.z, w: 0},
			max: point4d{x: c.bounds.max.x, y: c.bounds.max.y, z: c.bounds.max.z, w: 0},
		},
		data: data,
	}
}

func (c *hypercube) neighbors(p point4d) []point4d {
	var ps []point4d
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					q := point4d{p.x + dx, p.y + dy, p.z + dz, p.w + dw}
					if _, ok := c.data[q]; ok {
						ps = append(ps, q)
					}
				}
			}
		}
	}
	return ps
}

func (c *hypercube) evolve(fn evolveFn) {
	data := map[point4d]interface{}{}
	c.bounds.min.x--
	c.bounds.min.y--
	c.bounds.min.z--
	c.bounds.min.w--
	c.bounds.max.x++
	c.bounds.max.y++
	c.bounds.max.z++
	c.bounds.max.w++
	for w := c.bounds.min.w; w <= c.bounds.max.w; w++ {
		for z := c.bounds.min.z; z <= c.bounds.max.z; z++ {
			for y := c.bounds.min.y; y <= c.bounds.max.y; y++ {
				for x := c.bounds.min.x; x <= c.bounds.max.x; x++ {
					p := point4d{x, y, z, w}
					_, active := c.data[p]
					if fn(active, len(c.neighbors(p))) {
						data[p] = nil
					}
				}
			}
		}
	}
	c.data = data
}
