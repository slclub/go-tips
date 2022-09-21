package tips

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
	var b float64 = 23432.0
	if Any2Int64(b) != int64(23432) {
		t.Error("Any2Int64 not pass")
	}
}
