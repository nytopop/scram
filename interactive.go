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

			parsed := read(*result)
			res := Eval(parsed, &globalenv)
			out := String(res)

			fmt.Println(out)
		}
	}
}
