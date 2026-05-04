package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func echo(args []string, w io.Writer) {
	fmt.Fprintln(w, strings.Join(args, " "))
}

func main() {
	echo(os.Args[1:], os.Stdout)
}
