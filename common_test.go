package helper

import "testing"

func TestGetRootPath(t *testing.T) {
	p := GetRootPath()
	t.Log(p)
}

func TestAny2Int64(t *testing.T) {
	s := "012232"
	if si := Any2Int64(s); si != 12232 {
		t.Error("Any2Int64 not pass", si)
	}
	var a uint16 = 1232
	if si := Any2Int64(a); si != 1232 {
		t.Error("Any2Int64 not pass")
	}
}
