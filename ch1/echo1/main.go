// Echo1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"io"
	"os"
)

func echo(args []string, w io.Writer) {
	for i := 1; i < len(args); i++ {
		fmt.Fprintln(w, args[i])
	}
}

func main() {
	echo(os.Args, os.Stdout)
}
