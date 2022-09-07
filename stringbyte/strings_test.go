package stringbyte

import "testing"

func TestSubLeft(t *testing.T) {
	s := "sub.ex"
	if SubLeft(s, ".") != "sub" {
		t.Error("SubLeft not pass")
	}
	s = ""
	if SubLeft(s, ".") != s {
		t.Error("SubLeft not pass")
	}
}
