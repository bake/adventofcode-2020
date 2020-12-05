package passport

import (
	"bytes"
	"fmt"
)

type pair struct{ key, value []byte }

type Scanner struct {
	data   []byte
	pair   pair
	offset int
	err    error
}

func NewScanner(data []byte) *Scanner {
	return &Scanner{data: append(data, ' ')}
}

func (s Scanner) Error() error { return s.err }

func (s *Scanner) Next() bool {
	var n int
	var err error
	n, s.pair.key, err = s.until(":")
	if err != nil {
		return false
	}
	s.offset += n + 1
	n, s.pair.value, err = s.until(" \n")
	if err != nil {
		s.err = fmt.Errorf("value for %q missing", s.pair.key)
		return false
	}
	s.offset += n + 1
	return true
}

func (s *Scanner) until(delim string) (int, []byte, error) {
	n := bytes.IndexAny(s.data[s.offset:], delim)
	if n < 0 {
		return 0, nil, fmt.Errorf("end of input")
	}
	return n, s.data[s.offset : s.offset+n], nil
}
