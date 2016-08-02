package main

import (
	"bytes"
	"fmt"
)

func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := n; i > 0; i-- {
		if i%3 == 0 && i < n {
			buf.WriteString(",")
		}
		buf.WriteByte(s[n-i])
	}
	return buf.String()
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma2(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma1("12345678984564641278945646464613132"))
	fmt.Println(comma2("12345678984564641278945646464613132"))
}
