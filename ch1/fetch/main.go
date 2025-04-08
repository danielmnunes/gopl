// Fetch exibe o conteúdo encontrado em cada url especificado
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fecth: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fecth: reading %s: %v\n", url, err)
		}
		fmt.Printf("%s", b)
	}
}
