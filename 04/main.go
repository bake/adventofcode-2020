package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/bake/adventofcode-2020/04/passport"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	datas, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(datas))
	fmt.Println(part2(datas))
	return nil
}

func input(r io.Reader) ([][]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes.Split(data, []byte("\n\n")), nil
}

func part1(datas [][]byte) int {
	type port struct {
		BirthYear      int    `passport:"byr" validate:"required"`
		IssueYear      int    `passport:"iyr" validate:"required"`
		ExpirationYear int    `passport:"eyr" validate:"required"`
		Height         string `passport:"hgt" validate:"required"`
		HairColor      string `passport:"hcl" validate:"required"`
		EyeColor       string `passport:"ecl" validate:"required"`
		PassportID     string `passport:"pid" validate:"required"`
		CountryID      int    `passport:"cid"`
	}
	var num int
	for _, data := range datas {
		var p port
		if err := passport.Unmarshal(data, &p); err != nil {
			continue
		}
		if err := passport.Validate(p); err != nil {
			continue
		}
		num++
	}
	return num
}

func part2(datas [][]byte) int {
	type port struct {
		BirthYear      int      `passport:"byr" validate:"required,min=1920,max=2002"`
		IssueYear      int      `passport:"iyr" validate:"required,min=2010,max=2020"`
		ExpirationYear int      `passport:"eyr" validate:"required,min=2020,max=2030"`
		Height         height   `passport:"hgt" validate:"required"`
		HairColor      color    `passport:"hcl" validate:"required"`
		EyeColor       eyeColor `passport:"ecl" validate:"required"`
		PassportID     string   `passport:"pid" validate:"required,min=9,max=9"`
		CountryID      int      `passport:"cid"`
	}
	var num int
	for _, data := range datas {
		var p port
		if err := passport.Unmarshal(data, &p); err != nil {
			continue
		}
		if err := passport.Validate(p); err != nil {
			continue
		}
		num++
	}
	return num
}
