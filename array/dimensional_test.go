package array

import (
	"github.com/slclub/go-tips/logf"
	"strconv"
	"testing"
)

func TestDimensionPlugk(t *testing.T) {
	logf.LogOf(logf.New())

	data := []map[string]any{}
	for i := 0; i < 10; i++ {
		data = append(data, map[string]any{
			"ID":   i,
			"Name": "name" + strconv.Itoa(i),
		})
	}

	data = append(data, map[string]any{"Sex": 1})
	arr_int := []int{}
	DimensionPlugk(&arr_int, data, "ID")
	logf.Log().Print("get arr int:", arr_int)

	arr_string := []string{}
	DimensionPlugk(&arr_string, data, "Name")
	logf.Log().Print("get arr string:", arr_string)

	pusers := []any{}
	for i := 0; i < 10; i++ {
		u := user{
			ID:       i,
			Name:     "name" + strconv.Itoa(i),
			CreateAt: int64(i + 64),
			Sex:      int8(i % 2),
		}
		pusers = append(pusers, u)
		pusers = append(pusers, &u)
	}
	pusers = append(pusers, user1{})
	names := []string{}
	DimensionPlugk(&names, pusers, "Name")
	logf.Log().Print("get arr string of *struct:", names)

	ids := []int{}
	DimensionPlugk(&ids, pusers, "ID")
	logf.Log().Print("get arr int of *struct:", ids)
}

type user struct {
	ID       int
	Name     string
	CreateAt int64
	Sex      int8
}

type user1 struct {
}