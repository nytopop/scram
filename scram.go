// scream.go

package main

import "os"

func main() {
	if len(os.Args) < 2 {
		REPL("-> ")
	} else {
		// TODO: read args as files containing scram code
	}
}
