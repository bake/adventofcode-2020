package main

import (
	"strconv"
	"strings"
)

type parseFunc func() parseFunc

type RuleParser struct {
	data   string
	off    int
	name   string
	rules  []rule
	err    error
	parser parseFunc
}

func NewRuleParser(data string) *RuleParser {
	rp := &RuleParser{data: data}
	rp.parser = rp.parseName
	return rp
}

func (rp *RuleParser) parseName() parseFunc {
	n := rp.until(":")
	rp.name = rp.data[rp.off : rp.off+n]
	rp.off += n + 1
	return rp.parseRule
}

func (rp *RuleParser) parseRule() parseFunc {
	rp.rules = append(rp.rules, rule{
		rp.integer(),
		rp.integer(),
	})
	if rp.off >= len(rp.data) {
		return nil
	}
	return rp.parseRule
}

func (rp *RuleParser) Parse() bool {
	rp.parser = rp.parser()
	return rp.parser != nil
}

func (rp *RuleParser) integer() int {
	rp.off += rp.until("0123456789")
	n := rp.until("-: ")
	raw := rp.data[rp.off : rp.off+n]
	rp.off += n
	i, err := strconv.Atoi(raw)
	rp.err = err
	return i
}

func (rp *RuleParser) until(chars string) int {
	n := strings.IndexAny(rp.data[rp.off:], chars)
	if n < 0 {
		return len(rp.data) - rp.off
	}
	return n
}

func (rp *RuleParser) Err() error    { return rp.err }
func (rp *RuleParser) Name() string  { return rp.name }
func (rp *RuleParser) Rules() []rule { return rp.rules }
