package main

import (
	"testing"
)

const (
	number = "1234567890"
	en     = "hello world."
	zh     = "你好 世界。"
	jpn    = "こんにちは、世界."
	kr     = "안녕 하세요 세계."
)

var (
	tests = []struct{ in, want string }{
		{encode(number), number},
		{encode(en), en},
		{encode(zh), zh},
		{encode(jpn), jpn},
		{encode(kr), kr},
	}
)

func TestEncode(t *testing.T) {
	for _, test := range tests {
		if got := decode(test.in); got != test.want {
			t.Errorf("decode(%q) = %q, want %q", test.in, got, test.want)
		}
	}
}
