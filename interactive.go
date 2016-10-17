// interactive.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func ExecuteFiles(files []string) {
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			parsed := read(scanner.Text())
			res := Eval(parsed, &globalenv)
			out := String(res)

			fmt.Println(out)
		}
	}
}
