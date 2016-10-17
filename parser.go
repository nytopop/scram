// parser.go

package main

import (
	"strconv"
	"strings"
)

// Scram expression. All scram types implement this type.
type Scram interface{}

type Symbol string
type Number float64

// Read expression from string
func read(s string) Scram {
	tokens := Tokenize(s)
	return TokenTree(&tokens)
}

// Syntactic Analysis
func TokenTree(tokens *[]string) Scram {
	token := (*tokens)[0]
	*tokens = (*tokens)[1:]

	switch token {
	case "(":
		exp := make([]Scram, 0)
		for (*tokens)[0] != ")" {
			if i := TokenTree(tokens); i != Symbol("") {
				exp = append(exp, i)
			}
		}

		*tokens = (*tokens)[1:]
		return exp
	case ")":
		panic("Syntax Error: Unexpected )")
	default:
		if val, err := strconv.ParseFloat(token, 64); err == nil {
			return Number(val)
		} else {
			return Symbol(token)
		}
	}
}

// Lexical Analysis
func Tokenize(s string) []string {
	buf := strings.Replace(s, "(", "( ", -1)
	buf = strings.Replace(buf, ")", " )", -1)
	out := strings.Split(buf, " ")
	return out
}
