package tips

import (
	"errors"
	"testing"
)

func TestStrPos(t *testing.T) {
	str := "hello world go tips"

	i := StrPos(str, "wo")
	if i == -1 {
		t.Error("strings StrPos expect index >= 0; not -1")
	}

	i = StrPos(str, "v")
	if i != -1 {
		t.Error("strings StrPos expect -1 not index >= 0")
	}
	i = StrPos(str, "tip")
	if i == -1 {
		t.Error("strings StrPos expect index >= 0; not -1")
	}

	i = StrPos(str, "tipv")
	if i != -1 {
		t.Error("strings StrPos expect -1 not index >= 0")
	}

	i = StrPos(str, "")
	if i != 0 {
		t.Error("strings StrPos expect 0 not -1", i)
	}
}

func TestStrBegin(t *testing.T) {
	str := "hello world go tips"
	if StrBegin(str, "o") != "hell" {
		t.Error("strings StrPos expect hell")
	}
	if StrBegin(str, "tip") != "hello world go " {
		t.Error("strings StrPos expect 'hello world go '")
	}
	if StrBegin(str, "") != str {
		t.Error("strings StrPos expect '", str, "'")
	}
	if StrBegin(str, " ") != "hello" {
		t.Error("strings StrPos expect 'hello'")
	}
}

func TestStrEnd(t *testing.T) {
	str := "hello world go tips"
	if StrEnd(str, "o") != " world go tips" {
		t.Error("strings StrPos expect: world go tips")
	}
	if StrEnd(str, "tip") != "s" {
		t.Error("strings StrPos expect:s")
	}
	if StrEnd(str, "wo") != "rld go tips" {
		t.Error("strings StrPos expect:rld go tips")
	}
	if StrEnd(str, "") != "" {
		t.Error("strings StrPos expect:")
	}
	if StrEnd(str, " ") != "world go tips" {
		t.Error("strings StrPos expect:world go tips")
	}

}

func TestString(t *testing.T) {
	_trr := "error to string"
	trr := errors.New(_trr)
	if String(trr) != _trr {
		t.Fatal("TIPS.String  error to string")
	}
}
