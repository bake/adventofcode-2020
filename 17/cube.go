package main

type point3d struct{ x, y, z int }

type bounds3d struct{ min, max point3d }

type cube struct {
	bounds bounds3d
	data   map[point3d]interface{}
}

func newCube(r *rectangle) *cube {
	data := map[point3d]interface{}{}
	for k := range r.data {
		data[point3d{x: k.x, y: k.y, z: 0}] = nil
	}
	return &cube{
		bounds: bounds3d{
			min: point3d{x: r.bounds.min.x, y: r.bounds.min.y, z: 0},
			max: point3d{x: r.bounds.max.x, y: r.bounds.max.y, z: 0},
		},
		data: data,
	}
}

func (c *cube) neighbors(p point3d) []point3d {
	var ps []point3d
	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				q := point3d{p.x + dx, p.y + dy, p.z + dz}
				if _, ok := c.data[q]; ok {
					ps = append(ps, q)
				}
			}
		}
	}
	return ps
}

func (c *cube) evolve(fn evolveFn) {
	data := map[point3d]interface{}{}
	c.bounds.min.x--
	c.bounds.min.y--
	c.bounds.min.z--
	c.bounds.max.x++
	c.bounds.max.y++
	c.bounds.max.z++
	for z := c.bounds.min.z; z <= c.bounds.max.z; z++ {
		for y := c.bounds.min.y; y <= c.bounds.max.y; y++ {
			for x := c.bounds.min.x; x <= c.bounds.max.x; x++ {
				p := point3d{x, y, z}
				_, active := c.data[p]
				if fn(active, len(c.neighbors(p))) {
					data[p] = nil
				}
			}
		}
	}
	c.data = data
}
