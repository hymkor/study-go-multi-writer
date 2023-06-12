package multi_test

import (
	"io"
	"os"
	"testing"

	"github.com/hymkor/study-go-multi-writer"
)

func try(f func(int, io.Writer) string, n int) {
	devNull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer devNull.Close()

	f(n, devNull)
}

func BenchmarkMulti1(b *testing.B) {
	try(multi.Multi1, b.N)
}

func BenchmarkMulti2(b *testing.B) {
	try(multi.Multi2, b.N)
}

func BenchmarkMulti3(b *testing.B) {
	try(multi.Multi3, b.N)
}
