// interactive.go

package main

import (
	"fmt"
	"strings"

	"github.com/shavac/readline"
)

func String(v Scram) string {
	switch v := v.(type) {
	case []Scram:
		l := make([]string, len(v))
		for i, x := range v {
			l[i] = String(x)
		}
		return "(" + strings.Join(l, " ") + ")"
	default:
		return fmt.Sprint(v)
	}
}

func REPL(prompt string) {
	fmt.Println("a scheme is afoot!")

	for {
		switch result := readline.ReadLine(&prompt); true {
		case result == nil:
			println()
			break
		case *result != "":
			readline.AddHistory(*result)

			tokens := Tokenize(*result)
			exp := TokenTree(&tokens)
			val := Eval(exp, &globalenv)
			out := String(val)

			fmt.Println(out)
		}
	}
}

func ExecFiles(files []string) {
	for _, f := range files {
		s := ReadFile(f)
		exps := ReadExps(s)
		for _, raw := range exps {
			tokens := Tokenize(raw)
			exp := TokenTree(&tokens)
			val := Eval(exp, &globalenv)
			out := String(val)

			fmt.Println(out)
		}
	}
}
