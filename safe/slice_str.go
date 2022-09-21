package safe

import "bytes"

// ------------------------------------------------------
// slice string
type SliceString []string

func (this *SliceString) Append(val string) {
	*this = append(*this, val)
}

func (this *SliceString) AppendUnqiue(val string) {
	for i, n := 0, this.Len(); i < n; i++ {
		if (*this)[i] == val {
			return
		}
	}
	this.Append(val)
}
func (this *SliceString) AppendArr(arrs ...[]string) {
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
		*this = make(SliceString, ml)
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

func (this *SliceString) Del(val string) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			(*this)[i] = (*this)[n-1]
			(*this) = (*this)[:n-1]
			return i
		}
	}
	return -1
}

func (this *SliceString) DelKey(k int) string {
	n := len(*this)
	if k >= n {
		return ""
	}
	(*this)[k] = (*this)[n-1]
	(*this) = (*this)[:n-1]
	return (*this)[k]
}

func (this *SliceString) Reset() {
	*this = SliceString{}
}
func (this *SliceString) Len() int {
	return len(*this)
}

// find position of the val
// @return
// 		-1 	: not found
//  	>=0 : return the key of val
func (this *SliceString) In(val string) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i] == val {
			return i
		}
	}
	return -1
}

func (this *SliceString) Join(sep string) string {
	ml := len(*this)
	buf := new(bytes.Buffer)
	for i, v := range *this {
		buf.WriteString(v)
		if i < ml-1 {
			buf.WriteString(sep)
		}
	}
	return string(buf.Bytes())
}

func (this *SliceString) Range(fn func(i int, val string) bool) {
	for i, n := 0, len(*this); i < n; i++ {
		rtn := fn(i, (*this)[i])
		if rtn == false {
			return
		}
	}
}
