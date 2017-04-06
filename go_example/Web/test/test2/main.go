package main

import (
	"bytes"
	"fmt"
	"html"
	"os"
	"strings"

	"github.com/CodyGuo/logs"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	r, _ := os.Open("V3.5Ver_List.htm")
	doc, _ := goquery.NewDocumentFromReader(r)
	doc.Find(".MsoNormal").Each(func(i int, s *goquery.Selection) {
		style, _ := s.Find("span").Attr("style")
		if style == `font-size:18.0pt;line-height:150%;font-family:NSimSun;color:navy` {
			title := strings.TrimSpace(s.Text())
			fmt.Printf("title --> %s\n", title)
		}

		if style == `font-family:"Verdana","sans-serif";color:blue` {
			version := strings.TrimSpace(s.Text())
			fmt.Printf("version --> %s\n", version)
		}

		if style == `font-family:SimSun` {
			typeTitle := strings.TrimSpace(s.Text())
			fmt.Println(typeTitle)
		}

		if style == `font-family:"Verdana","sans-serif"` {
			data := strings.Replace(s.Text(), "\n", "", -1)
			fmt.Println(data)
		}

	})
}

func encode_examples() {
	str := "1234567890. 你好 世界。 hello wrold. こんにちは、世界. 안녕 하세요 세계."
	info := encode(str)
	logs.Notice(info)
	logs.Noticef("info --> %s", decode(info))
}

func decode(str string) string {
	return html.UnescapeString(str)
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
