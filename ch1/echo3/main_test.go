package main

import (
	"io"
	"testing"
)

var args = []string{"hello", "world", "foo", "bar", "baz"}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo(args, io.Discard)
	}
}
