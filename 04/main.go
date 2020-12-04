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

type port struct {
	BirthYear      int    `passport:"required,byr"`
	IssueYear      int    `passport:"required,iyr"`
	ExpirationYear int    `passport:"required,eyr"`
	Height         string `passport:"required,hgt"`
	HairColor      string `passport:"required,hcl"`
	EyeColor       string `passport:"required,ecl"`
	// Height and Passport ID might get mixed up, so this needs to be a string.
	PassportID string `passport:"required,pid"`
	CountryID  int    `passport:"cid"`
}

func run() error {
	datas, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(datas))
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
	var num int
	for _, data := range datas {
		var p port
		err := passport.Unmarshal(data, &p)
		if err != nil {
			continue
		}
		num++
	}
	return num
}
