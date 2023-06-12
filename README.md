io.MultiWriter は遅いから、使わなくてもよい時は使わない方がよいというだけのリポート
=====================

string とファイル（ここでは /dev/null）へ同じものを出力する

```main.go
package multi

import (
    "bufio"
    "io"
    "strings"
)

// Multi1 - strings.Builder で string のインスタンスを作ったあと、それをファイルにも出力する
func Multi1(n int, w io.Writer) string {
    var buffer strings.Builder
    for i := 0; i < n; i++ {
        buffer.WriteByte(byte(i))
    }
    s := buffer.String()
    io.WriteString(w, s)
    return s
}

// Multi2 - MultiWriter で、strings.Builder とファイルに同時出力
func Multi2(n int, w io.Writer) string {
    var buffer strings.Builder
    ww := io.MultiWriter(w, &buffer)
    for i := 0; i < n; i++ {
        ww.Write([]byte{byte(i)})
    }
    return buffer.String()
}

// Multi3 - Multi2 と同じだが、bufio でオーバーヘッドを軽減
func Multi3(n int, w io.Writer) string {
    var buffer strings.Builder
    bw := bufio.NewWriter(io.MultiWriter(w, &buffer))
    for i := 0; i < n; i++ {
        bw.WriteByte(byte(i))
    }
    bw.Flush()
    return buffer.String()
}
```

```main_test.go
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
```

### 「go test -bench . -benchmem」の結果

``` go test -bench . -benchmem |
goos: windows
goarch: amd64
pkg: github.com/hymkor/study-go-multi-writer
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkMulti1-4   	355637443	         3.080 ns/op	       5 B/op	       0 allocs/op
BenchmarkMulti2-4   	  734821	      1609 ns/op	       6 B/op	       1 allocs/op
BenchmarkMulti3-4   	272213160	         4.154 ns/op	       5 B/op	       0 allocs/op
PASS
ok  	github.com/hymkor/study-go-multi-writer	5.670s
```
