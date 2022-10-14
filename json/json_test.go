package json

import (
	"github.com/slclub/go-tips/logf"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	logf.LogOf(logf.New())
	u := user{
		Uid:  1029320192,
		Name: "halod kit",
	}

	b, err := Marshal(u)
	if err != nil {
		t.Error("JSON Marshal error: ", err)
	}
	tu := map[string]any{}
	Unmarshal(b, &tu)
	logf.Log().Print("unmarshal", tu)

	Decode(b, &tu)
	logf.Log().Print("decode", tu)
}

type user struct {
	Uid  int
	Name string
}
