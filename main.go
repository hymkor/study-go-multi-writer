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
