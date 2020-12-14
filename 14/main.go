package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ps, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(ps))
	fmt.Println(part2(ps))
	return nil
}

type program struct {
	mask   mask
	memory []write
}

type write struct {
	address, value int
}

func input(r io.Reader) ([]program, error) {
	var ps []program
	s := bufio.NewScanner(r)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "mask = ") {
			ps = append(ps, program{
				mask: mask(strings.TrimPrefix(s.Text(), "mask = ")),
			})
			continue
		}
		var k, v int
		if _, err := fmt.Sscanf(s.Text(), "mem[%d] = %d", &k, &v); err != nil {
			return nil, err
		}
		ps[len(ps)-1].memory = append(ps[len(ps)-1].memory, write{k, v})
	}
	return ps, s.Err()
}

func part1(ps []program) int {
	mem := map[int]int{}
	for _, p := range ps {
		for _, w := range p.memory {
			mem[w.address] = w.value
			for i, b := range p.mask {
				i = 35 - i
				switch b {
				case '0':
					mem[w.address] &= ^(1 << i)
				case '1':
					mem[w.address] |= (1 << i)
				}
			}
		}
	}
	var sum int
	for _, v := range mem {
		sum += v
	}
	return sum
}

type mask string

// apply the mask on a given address by expanding all Xs.
func (m mask) apply(address string) []string {
	num := strings.Count(string(m), "X")
	max := int(math.Pow(2, float64(num)))
	var as []string
	for i := 0; i < max; i++ {
		n := []byte(address)
		bs := fmt.Sprintf("%0*b", bits.Len(uint(max-1)), i)
		for i, r := range m {
			switch r {
			case '0':
			case '1':
				n[i] = '1'
			case 'X':
				n[i], bs = bs[0], bs[1:]
			}
		}
		as = append(as, string(n))
	}
	return as
}

func part2(ps []program) int {
	mem := map[int64]int{}
	for _, p := range ps {
		for _, w := range p.memory {
			address := fmt.Sprintf("%0*b", len(p.mask), w.address)
			for _, a := range p.mask.apply(address) {
				a, err := strconv.ParseInt(a, 2, 64)
				if err != nil {
					log.Fatal(err)
					return 0
				}
				mem[a] = w.value
			}
		}
	}
	var sum int
	for _, v := range mem {
		sum += v
	}
	return sum
}
