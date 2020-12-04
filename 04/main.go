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
		BirthYear      int    `passport:"required,byr"`
		IssueYear      int    `passport:"required,iyr"`
		ExpirationYear int    `passport:"required,eyr"`
		Height         string `passport:"required,hgt"`
		HairColor      string `passport:"required,hcl"`
		EyeColor       string `passport:"required,ecl"`
		PassportID     string `passport:"required,pid"`
		CountryID      int    `passport:"cid"`
	}
	var num int
	for _, data := range datas {
		var p port
		if err := passport.Unmarshal(data, &p); err != nil {
			continue
		}
		num++
	}
	return num
}

func part2(datas [][]byte) int {
	type port struct {
		BirthYear      int      `passport:"required,byr" validate:"min=1920,max=2002"`
		IssueYear      int      `passport:"required,iyr" validate:"min=2010,max=2020"`
		ExpirationYear int      `passport:"required,eyr" validate:"min=2020,max=2030"`
		Height         height   `passport:"required,hgt"`
		HairColor      color    `passport:"required,hcl"`
		EyeColor       eyeColor `passport:"required,ecl"`
		PassportID     string   `passport:"required,pid" validate:"min=9,max=9"`
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
