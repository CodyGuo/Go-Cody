package main

import (
	"testing"
)

func TestStar(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak) = false`)
	}

	if !IsPalindrome("te1et") {
		t.Error(`IsPalindrome(test) = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("hello") {
		t.Error(`IsPalindrome(helloolleh) = true`)
	}
}
