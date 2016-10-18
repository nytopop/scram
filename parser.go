// parser.go

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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

// Parse string into expressions
func ReadExps(in string) []string {
	exps := []string{}
	var tmp string
	par := 0

	for _, r := range in {
		c := fmt.Sprintf("%c", r)
		tmp += c
		if c == "(" {
			par++
		} else if c == ")" {
			par--
			if par == 0 {
				exps = append(exps, tmp)
				tmp = ""
			}
		}
	}

	return exps
}

// Read file into string
func ReadFile(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	s := string(buf.Bytes())
	clean := strings.Replace(s, "\n", "", -1)
	clean = strings.Replace(clean, "\r", "", -1)
	clean = strings.Replace(clean, "\t", "", -1)

	return clean
}
