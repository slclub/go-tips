package safe

// ------------------------------------------------------
// slice Value

type Value interface {
	Value() int64
}

type SliceValue []Value

func (this *SliceValue) Append(val Value) {
	*this = append(*this, val)
}

func (this *SliceValue) AppendUnqiue(val Value) {
	for i, n := 0, this.Len(); i < n; i++ {
		if (*this)[i].Value() == val.Value() {
			(*this)[i] = val // 覆盖
			return
		}
	}
	this.Append(val)
}
func (this *SliceValue) AppendArr(arrs ...[]Value) {
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
		*this = make(SliceValue, ml)
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

func (this *SliceValue) Del(val Value) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i].Value() == val.Value() {
			(*this)[i] = (*this)[n-1]
			(*this) = (*this)[:n-1]
			return i
		}
	}
	return -1
}

func (this *SliceValue) DelKey(k int) Value {
	n := len(*this)
	if k >= n {
		return nil
	}
	(*this)[k] = (*this)[n-1]
	(*this) = (*this)[:n-1]
	return (*this)[k]
}

func (this *SliceValue) Reset() {
	*this = SliceValue{}
}
func (this *SliceValue) Len() int {
	return len(*this)
}

// find position of the val
// @return
// 		-1 	: not found
//  	>=0 : return the key of val
func (this *SliceValue) In(val Value) int {
	for i, n := 0, len(*this); i < n; i++ {
		if (*this)[i].Value() == val.Value() {
			return i
		}
	}
	return -1
}

func (this *SliceValue) Range(fn func(i int, val Value) bool) {
	for i, n := 0, len(*this); i < n; i++ {
		rtn := fn(i, (*this)[i])
		if rtn == false {
			return
		}
	}
}
