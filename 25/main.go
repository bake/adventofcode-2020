package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	card, door, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(card, door))
	return nil
}

type direction string

func input(r io.Reader) (int, int, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, 0, err
	}
	keys := bytes.Split(data, []byte{'\n'})
	card, err := strconv.Atoi(string(keys[0]))
	if err != nil {
		return 0, 0, err
	}
	door, err := strconv.Atoi(string(keys[1]))
	if err != nil {
		return 0, 0, err
	}
	return card, door, nil
}

func transform(subject, loops int) int {
	value := 1
	for i := 0; i < loops; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}

func loops(subject, public int) (i int) {
	value := 1
	for i = 0; value != public; i++ {
		value *= subject
		value %= 20201227
	}
	return i
}

func part1(cardPK, doorPK int) int {
	cardLoops := loops(7, cardPK)
	return transform(doorPK, cardLoops)
}
