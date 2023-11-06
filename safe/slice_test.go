package safe

import (
	"github.com/slclub/go-tips/logf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceInt(t *testing.T) {
	s1 := []int{1, 2, 3}
	ss1 := SliceInt(s1)
	ss1.Append(4)
	assert.True(t, len(ss1) == 4)
	assert.True(t, ss1.In(3) > 0)

	ss1.AppendArr([]int{5, 6}, []int{10, 11})
	logf.Print("SliceInt.Join:=", ss1.Join(","))

	ss1.Range(func(i, v int) bool {
		if v%2 == 0 {
			return false
		}
		return true
	})

	ss1.Del(1)
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(4)
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(20)
	assert.True(t, ss1.Len() == 8)

	assert.True(t, len(ss1[:3]) == 3)

	ss1.Reset()
	assert.True(t, len(ss1) == 0)
}

func TestSliceInt32(t *testing.T) {
	s1 := []int32{1, 2, 3}
	ss1 := SliceInt32(s1)
	ss1.Append(4)
	assert.True(t, len(ss1) == 4)
	assert.True(t, ss1.In(3) > 0)

	ss1.AppendArr([]int32{5, 6}, []int32{10, 11})
	logf.Print("SliceInt64.Join:=", ss1.Join(","))

	ss1.Range(func(i int, v int32) bool {
		if v%2 == 0 {
			return false
		}
		return true
	})

	ss1.Del(1)
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(4)
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(20)
	assert.True(t, ss1.Len() == 8)

	assert.True(t, len(ss1[:3]) == 3)

	ss1.Reset()
	assert.True(t, len(ss1) == 0)
}
func TestSliceInt64(t *testing.T) {
	s1 := []int64{1, 2, 3}
	ss1 := SliceInt64(s1)
	ss1.Append(4)
	assert.True(t, len(ss1) == 4)
	assert.True(t, ss1.In(3) > 0)

	ss1.AppendArr([]int64{5, 6}, []int64{10, 11})
	logf.Print("SliceInt64.Join:=", ss1.Join(","))

	ss1.Range(func(i int, v int64) bool {
		if v%2 == 0 {
			return false
		}
		return true
	})

	ss1.Del(1)
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(4)
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(20)
	assert.True(t, ss1.Len() == 8)

	assert.True(t, len(ss1[:3]) == 3)

	ss1.Reset()
	assert.True(t, len(ss1) == 0)
}

func TestSliceString(t *testing.T) {
	s1 := []string{"1", "2", "3"}
	ss1 := SliceString(s1)
	ss1.Append("4")
	assert.True(t, len(ss1) == 4)
	assert.True(t, ss1.In("3") > 0)

	ss1.AppendArr([]string{"51", "6"}, []string{"a", "bb"})
	logf.Print("SliceInt64.Join:=", ss1.Join(","))

	ss1.Range(func(i int, v string) bool {
		if len(v) >= 2 {
			return false
		}
		return true
	})

	ss1.Del("1")
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue("4")
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue("20")
	assert.True(t, ss1.Len() == 8)

	assert.True(t, len(ss1[:3]) == 3)

	ss1.Reset()
	assert.True(t, len(ss1) == 0)

	_ = append(ss1, s1...)
}

type valueInt int

func (this valueInt) Value() int64 {
	return int64(this)
}

func TestSliceValue(t *testing.T) {
	ss1 := SliceValue{valueInt(1), valueInt(2), valueInt(3)}
	ss1.Append(valueInt(4))
	assert.True(t, len(ss1) == 4)
	assert.True(t, ss1.In(valueInt(3)) > 0)

	ss1.AppendArr(SliceValue{valueInt(5), valueInt(6)}, SliceValue{valueInt(10), valueInt(11)})

	ss1.Range(func(i int, v Value) bool {
		if v.Value() >= 6 {
			return false
		}
		return true
	})

	ss1.Del(valueInt(1))
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(valueInt(10))
	assert.True(t, ss1.Len() == 7)
	ss1.AppendUnqiue(valueInt(20))
	assert.True(t, ss1.Len() == 8)

	assert.True(t, len(ss1[:3]) == 3)
	logf.Print("SliceValue", ss1)
	ss1.Reset()
	assert.True(t, len(ss1) == 0)

}
