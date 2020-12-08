package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	insts, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(insts))
	fmt.Println(part2(insts))
	return nil
}

func input(r io.Reader) (*program, error) {
	s := bufio.NewScanner(r)
	var insts []instruction
	for s.Scan() {
		var name string
		var mod byte
		var argument int
		_, err := fmt.Sscanf(s.Text(), "%s %c%d", &name, &mod, &argument)
		if err != nil {
			continue
		}
		if mod == '-' {
			argument *= -1
		}
		insts = append(insts, instruction{name, argument})
	}
	return newProgram(insts), nil
}

func part1(p *program) int {
	p.run()
	return p.acc
}

func part2(p *program) int {
	for i, ins := range p.insts {
		if ins.name == "jmp" {
			q := p.modify(i, "nop")
			if err := q.run(); err == nil {
				return q.acc
			}
		}
		if ins.name == "nop" {
			q := p.modify(i, "jmp")
			if err := q.run(); err == nil {
				return q.acc
			}
		}
	}
	return 0
}

type instruction struct {
	name     string
	argument int
}

type program struct {
	ptr, acc int
	insts    []instruction
	err      error
	seen     map[int]interface{}
}

func newProgram(insts []instruction) *program {
	return &program{
		insts: insts,
		seen:  map[int]interface{}{},
	}
}

func (p *program) run() error {
	for p.next() {
	}
	return p.err
}

// next executes the current instruction and return if at least one is left.
func (p *program) next() bool {
	if p.ptr >= len(p.insts) {
		return false
	}
	if _, ok := p.seen[p.ptr]; ok {
		p.err = fmt.Errorf("unexpected loop")
		return false
	}
	p.seen[p.ptr] = nil
	ins := p.insts[p.ptr]
	switch ins.name {
	case "nop":
		p.ptr++
		return true
	case "jmp":
		p.ptr += ins.argument
		return true
	case "acc":
		p.ptr++
		p.acc += ins.argument
		return true
	}
	return false
}

// Change the instruction at a given position and create a new program.
func (p *program) modify(ptr int, name string) *program {
	dsts := make([]instruction, len(p.insts))
	copy(dsts, p.insts)
	dsts[ptr].name = name
	return newProgram(dsts)
}
