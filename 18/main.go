package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
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
	eqs, err := input(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(part1(eqs))
	fmt.Println(part2(eqs))
	return nil
}

func input(r io.Reader) ([]string, error) {
	var eqs []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		eqs = append(eqs, s.Text())
	}
	return eqs, s.Err()
}

func part1(eqs []string) int {
	var sum int
	for _, eq := range eqs {
		_, val := eval(eq)
		sum += val
	}
	return sum
}

type opFn func(int) int

func addOp(x int) func(int) int { return func(y int) int { return x + y } }
func mulOp(x int) func(int) int { return func(y int) int { return x * y } }

func eval(eq string) (n int, val int) {
	var head byte
	var op opFn = addOp(0)
	for len(eq) > 0 {
		head, eq, n = eq[0], eq[1:], n+1
		switch head {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lav, _ := strconv.Atoi(string(head))
			val = op(lav)
		case '+':
			op = addOp(val)
		case '*':
			op = mulOp(val)
		case '(':
			m, lav := eval(eq)
			n, val, eq = n+m, op(lav), eq[m:]
		case ')':
			return n, val
		}
	}
	return n, val
}

// I thought about using RPN or adding parentheses around + but found a small
// example on using Gos parser to evaluate mathematical expressions[1]. To
// change the precedence I simply swapped `+` for `/` and `*` for `-`.
//
// [1]: https://thorstenball.com/blog/2016/11/16/putting-eval-in-go/
func part2(eqs []string) int {
	var sum int
	for _, eq := range eqs {
		sum += eval2(eq)
	}
	return sum
}

func eval2(eq string) int {
	eq = strings.ReplaceAll(eq, "+", "/")
	eq = strings.ReplaceAll(eq, "*", "-")
	exp, _ := parser.ParseExpr(eq)
	return eval2exp(exp)
}

func eval2exp(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return eval2binary(exp)
	case *ast.ParenExpr:
		return eval2exp(exp.X)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}
	return 0
}

func eval2binary(exp *ast.BinaryExpr) int {
	left := eval2exp(exp.X)
	right := eval2exp(exp.Y)
	switch exp.Op {
	case token.QUO:
		return left + right
	case token.SUB:
		return left * right
	}
	return 0
}
