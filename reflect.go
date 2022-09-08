package tips

import (
	"reflect"
	"runtime"
)

// Get function name.
func FUNC_NAME(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Check the param is or not an function.
func IS_FUNC(fn interface{}) bool {

	if reflect.ValueOf(fn).Kind() == reflect.Func {
		return true
	}

	return false
}

// check type struct
// 判断是结构体
func IS_STRUCT(stu interface{}) bool {
	if reflect.ValueOf(stu).Kind() == reflect.Struct {
		return true
	}
	return false
}

// check is slice
func IS_SLICE(s interface{}) bool {
	if reflect.ValueOf(s).Kind() == reflect.Slice {
		return true
	}
	return false
}
