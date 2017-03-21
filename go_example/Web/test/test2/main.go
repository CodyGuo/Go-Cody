package main

import (
	"bytes"
	"fmt"
	"html"

	"github.com/CodyGuo/logs"
)

func main() {
	str := "1234567890. 你好 世界。 hello wrold. こんにちは、世界. 안녕 하세요 세계."
	info := encode(str)
	logs.Notice(info)
	logs.Noticef("info --> %s", html.UnescapeString(info))
}

func encode(str string) string {
	if len(str) == 0 {
		return str
	}
	var buf bytes.Buffer
	for _, r := range str {
		// \u 表示Unicode编码, 中日韩统一表意文字(CJK Unified Ideographs) [\u2E80-\uFE4F]
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
