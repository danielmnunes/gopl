package main

import (
	"fmt"
	"io"
	"os"
)

func echo(args []string, w io.Writer) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = "_"
	}
	fmt.Fprintln(w, s)
}

func main() {
	echo(os.Args[1:], os.Stdout)
}
