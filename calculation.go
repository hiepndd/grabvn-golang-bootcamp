package main

import (
	"bytes"
	"strconv"
	"unicode"
)

type comptometer interface {
	calculate(expression string) (int, error)
}

type repl struct {
	result int
}

func (repl *repl) calculate(expression string) (int, error) {
	var buf bytes.Buffer
	// Remove all white spaces
	for _, r := range expression {
		if unicode.IsNumber(r) || r == '+' || r == '-' || r == '/' || r == '*' {
			buf.WriteRune(r)
		}
		if unicode.IsLetter(r) {
			return 0, ErrExpressionContainAlphabet
		}
	}

	if buf.Len() == 0 {
		return 0, ErrExpressionBlank
	}

	expression = buf.String()
	l := len(expression)
	i := 0

	for i < l && unicode.IsNumber(rune(expression[i])) {
		i++
	}
	cur, _ := strconv.Atoi(expression[:i])

	res := 0
	for i < l {
		operators := expression[i]
		i++
		start := i
		for i < l && unicode.IsNumber(rune(expression[i])) {
			i++
		}
		a, _ := strconv.Atoi(expression[start:i])

		switch operators {
		case '+':
			res += cur
			cur = a
		case '-':
			res += cur
			cur = -1 * a
		case '*':
			cur *= a
		case '/':
			cur /= a
		}
	}
	repl.result = res + cur
	return res + cur, nil
}
