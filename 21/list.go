package main

// A list contains a slice of strings with additional functionality. This is a
// pretty bad way of handling this kind of data, but it's more fun than using a
// map inside a map. Again.
type list []string

func (l list) contains(v string) bool {
	for _, w := range l {
		if v == w {
			return true
		}
	}
	return false
}

func (l list) unique() list {
	m := map[string]struct{}{}
	for _, v := range l {
		m[v] = struct{}{}
	}
	var j list
	for v := range m {
		j = append(j, v)
	}
	return j
}

func (l list) diff(j list) list {
	var k list
	for _, v := range l {
		if !j.contains(v) {
			k = append(k, v)
		}
	}
	for _, v := range j {
		if !l.contains(v) {
			k = append(k, v)
		}
	}
	return k
}

func (l list) delete(v string) list {
	var j list
	for _, w := range l {
		if w == v {
			continue
		}
		j = append(j, w)
	}
	return j
}

func (l list) intersect(j list) list {
	var k list
	for _, v := range l {
		if j.contains(v) {
			k = append(k, v)
		}
	}
	return k
}
