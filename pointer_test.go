package helper

import "testing"

func TestIsNil(t *testing.T) {
	b := new(string)
	c := "a"
	b = nil
	if IsNil(b) == false {
		t.Error("b should be nil")
	}
	if IsNil(c) == true {
		t.Error("c is not an empty string")
	}
}
