package safe

import (
	"bytes"
	"strconv"
)

// ------------------------------------------------------
// slice int
type SliceInt []int

func (this *SliceInt) Append(val int) {
	*this = append(*this, val)
}

func (this *SliceInt) AppendUnqiue(val int) {
	for i, n := 0, this.Len(); i < n; i++ {
		if (*this)[i] == val {
			return
		}
	}
	this.Append(val)
}

func (this *SliceInt) AppendArr(arrs ...[]int) {
	if arrs == nil {
		return
	}
	ml := len(*this)
	index := ml - 1
	cap_len := cap(*this)
	for _, arr := range arrs {
		ml += len(arr)
	}
	// Is need malloc memory
	if cap_len < ml {
		t := *this
		*this = make([]int, ml)
		copy(*this, t)
	}
	// combined
	for _, arr := range arrs {
		if arr == nil {
			continue
		}
		for i, _ := range arr {
			//this.Append(arr[i])
			index++
			(*this)[index] = arr[i]
		}
	}
}

func (this *SliceInt) Del(val int) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			(*this)[i] = (*this)[n-1]
			(*this) = (*this)[:n-1]
			return i
		}
	}
	return -1
}

func (this *SliceInt) DelKey(k int) int {
	n := len(*this)
	if k >= n {
		return 0
	}
	o := (*this)[k]
	(*this)[k] = (*this)[n-1]
	(*this) = (*this)[:n-1]
	return o
}

func (this *SliceInt) Reset() {
	*this = []int{}
}

func (this *SliceInt) Len() int {
	return len(*this)
}

// find position of the val
// @return
// 		-1 	: not found
//  	>=0 : return the key of val
func (this *SliceInt) In(val int) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			return i
		}
	}
	return -1
}

func (this *SliceInt) Join(sep string) string {
	ml := len(*this)
	buf := new(bytes.Buffer)
	for i, v := range *this {
		vs := strconv.Itoa(v)
		buf.WriteString(vs)
		if i < ml-1 {
			buf.WriteString(sep)
		}
	}
	return string(buf.Bytes())
}

func (this *SliceInt) Range(fn func(i, val int) bool) {
	for i, n := 0, len(*this); i < n; i++ {
		rtn := fn(i, (*this)[i])
		if rtn == false {
			return
		}
	}
}

// ------------------------------------------------------
// slice int32
type SliceInt32 []int32

func (this *SliceInt32) Append(val int32) {
	*this = append(*this, val)
}

func (this *SliceInt32) AppendUnqiue(val int32) {
	for i, n := 0, this.Len(); i < n; i++ {
		if (*this)[i] == val {
			return
		}
	}
	this.Append(val)
}

func (this *SliceInt32) AppendArr(arrs ...[]int32) {
	if arrs == nil {
		return
	}
	ml := len(*this)
	index := ml - 1
	cap_len := cap(*this)
	for _, arr := range arrs {
		ml += len(arr)
	}
	// Is need malloc memory
	if cap_len < ml {
		t := *this
		*this = make(SliceInt32, ml)
		copy(*this, t)
	}
	// combined
	for _, arr := range arrs {
		if arr == nil {
			continue
		}
		for i, _ := range arr {
			//this.Append(arr[i])
			index++
			(*this)[index] = arr[i]
		}
	}
}

func (this *SliceInt32) Del(val int32) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			(*this)[i] = (*this)[n-1]
			(*this) = (*this)[:n-1]
			return i
		}
	}
	return -1
}

func (this *SliceInt32) DelKey(k int) int32 {
	n := len(*this)
	if k >= n {
		return 0
	}
	o := (*this)[k]
	(*this)[k] = (*this)[n-1]
	(*this) = (*this)[:n-1]
	return o
}

func (this *SliceInt32) Reset() {
	*this = SliceInt32{}
}
func (this *SliceInt32) Len() int {
	return len(*this)
}

// find position of the val
// @return
// 		-1 	: not found
//  	>=0 : return the key of val
func (this *SliceInt32) In(val int32) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			return i
		}
	}
	return -1
}

func (this *SliceInt32) Join(sep string) string {
	ml := len(*this)
	buf := new(bytes.Buffer)
	for i, v := range *this {
		vs := strconv.FormatInt(int64(v), 10)
		buf.WriteString(vs)
		if i < ml-1 {
			buf.WriteString(sep)
		}
	}
	return string(buf.Bytes())
}

func (this *SliceInt32) Range(fn func(i int, val int32) bool) {
	for i, n := 0, len(*this); i < n; i++ {
		rtn := fn(i, (*this)[i])
		if rtn == false {
			return
		}
	}
}

// ------------------------------------------------------
// slice int64
type SliceInt64 []int64

func (this *SliceInt64) Append(val int64) {
	*this = append(*this, val)
}
func (this *SliceInt64) AppendUnqiue(val int64) {
	for i, n := 0, this.Len(); i < n; i++ {
		if (*this)[i] == val {
			return
		}
	}
	this.Append(val)
}

func (this *SliceInt64) AppendArr(arrs ...[]int64) {
	if arrs == nil {
		return
	}
	ml := len(*this)
	index := ml - 1
	cap_len := cap(*this)
	for _, arr := range arrs {
		ml += len(arr)
	}
	// Is need malloc memory
	if cap_len < ml {
		t := *this
		*this = make(SliceInt64, ml)
		copy(*this, t)
	}
	// combined
	for _, arr := range arrs {
		if arr == nil {
			continue
		}
		for i, _ := range arr {
			//this.Append(arr[i])
			index++
			(*this)[index] = arr[i]
		}
	}
}

func (this *SliceInt64) Del(val int64) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			(*this)[i] = (*this)[n-1]
			(*this) = (*this)[:n-1]
			return i
		}
	}
	return -1
}

func (this *SliceInt64) DelKey(k int) int64 {
	n := len(*this)
	if k >= n {
		return 0
	}
	o := (*this)[k]
	(*this)[k] = (*this)[n-1]
	(*this) = (*this)[:n-1]
	return o
}

func (this *SliceInt64) Reset() {
	*this = SliceInt64{}
}
func (this *SliceInt64) Len() int {
	return len(*this)
}

// find position of the val
// @return
// 		-1 	: not found
//  	>=0 : return the key of val
func (this *SliceInt64) In(val int64) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			return i
		}
	}
	return -1
}

func (this *SliceInt64) Join(sep string) string {
	ml := len(*this)
	buf := new(bytes.Buffer)
	for i, v := range *this {
		vs := strconv.FormatInt(v, 10)
		buf.WriteString(vs)
		if i < ml-1 {
			buf.WriteString(sep)
		}
	}
	return string(buf.Bytes())
}

func (this *SliceInt64) Range(fn func(i int, val int64) bool) {
	for i, n := 0, len(*this); i < n; i++ {
		rtn := fn(i, (*this)[i])
		if rtn == false {
			return
		}
	}
}
