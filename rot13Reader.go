package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)

	if err != nil {
		return
	}

	// 1文字ずつ
	for i := 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
			b[i] += 13 // a~mの場合は13右へ

		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'x') {
			b[i] -= 13 // n~zの場合は左へ
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}    // 変換？？？
	io.Copy(os.Stdout, &r) // stdout に rを出力
}
