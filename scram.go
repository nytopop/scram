// scream.go

package main

import "os"

func main() {
	if len(os.Args) < 2 {
		REPL("-> ")
	} else {
		ExecFiles(os.Args[1:])
		REPL("-> ")
	}
}
