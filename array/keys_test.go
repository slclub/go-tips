package array

import "testing"

func TestMapKeys(t *testing.T) {
	data := map[string]int{"No": 1, "Yes": 3}
	data2 := map[string]string{}
	var data3 error

	if arr := MapKeys(data); len(arr) != 2 || arr[1] != "Yes" {
		t.Error("MapKeys check data error")
	}

	if arr := MapKeys(data2); arr != nil {
		t.Error("MapKeys return empty is not nil")
	}

	if arr := MapKeys(data3); arr != nil {
		t.Error("MapKeys argument is not supported empty")
	}
}
