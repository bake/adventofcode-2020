package main

import (
	"strconv"
	"strings"
)

type list struct {
	current *node
	size    int
}

func (l *list) copy() *list {
	j := &list{}
	n := l.current
	for i := 0; i < l.size; i++ {
		j.add(&node{label: n.label})
		n = n.next
	}
	j.next()
	return j
}

func (l *list) String() string {
	var xs []string
	n := l.current
	for i := 0; i < l.size; i++ {
		x := n.String()
		if n.label == l.current.label {
			x = "(" + x + ")"
		}
		xs = append(xs, x)
		n = n.next
	}
	return strings.Join(xs, ", ")
}

func (l *list) add(n *node) *node {
	l.size++
	if l.current == nil {
		l.current = n
		l.current.prev = l.current
		l.current.next = l.current
		return l.current
	}
	n.prev = l.current
	n.next = l.current.next
	l.current.next.prev = n
	l.current.next = n
	l.current = n
	return l.current
}

// next moved the current node by one position.
func (l *list) next() *node {
	l.current = l.current.next
	return l.current
}

// take returns the current node out of the list.
func (l *list) take() *node {
	n := l.current
	l.current.next.prev = n.prev
	l.current.prev.next = n.next
	l.current = n.next
	l.size--
	return n
}

// find returns a node with a given label.
func (l *list) find(label int) *node {
	n := l.current
	for i := 0; i < l.size; i++ {
		if n.label == label {
			return n
		}
		n = n.next
	}
	return nil
}

func (l *list) end() *node {
	for i := 0; i < l.size-1; i++ {
		l.next()
	}
	return l.current
}

func (l *list) min() *node {
	var min *node
	n := l.current
	for i := 0; i < l.size; i++ {
		if min == nil || n.label < min.label {
			min = n
		}
		n = n.next
	}
	return min
}

func (l *list) max() *node {
	var min *node
	n := l.current
	for i := 0; i < l.size; i++ {
		if min == nil || n.label > min.label {
			min = n
		}
		n = n.next
	}
	return min
}

type node struct {
	prev, next *node
	label      int
}

func (n *node) String() string {
	return strconv.Itoa(n.label)
}
