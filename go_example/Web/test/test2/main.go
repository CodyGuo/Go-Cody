package main

import (
	"bytes"
	"fmt"
	"html"

	"github.com/CodyGuo/logs"
)

func main() {
	info := encode("hello 世界.")
	logs.Notice(info)
	logs.Noticef("info --> %s", html.UnescapeString(info))
}

func encode(str string) string {
	if len(str) == 0 {
		return str
	}
	var buf bytes.Buffer
	for _, r := range str {
		if r >= '\u2E80' && r <= '\uFE4F' {
			buf.Write([]byte("&#"))
			buf.Write([]byte(fmt.Sprintf("%d", r)))
			buf.Write([]byte(";"))
		} else {
			buf.Write([]byte(fmt.Sprintf("%c", r)))
		}
	}
	return buf.String()
}
