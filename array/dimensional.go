package array

import (
	"github.com/slclub/go-tips/convtype"
	"reflect"
)

// 多维数组

// 从二维数组中某一个元素取值 做为一个新一维数组

func DimensionPlugk(target any, data any, field any) {
	switch data_arr := data.(type) {
	case []map[string]any, []map[int]any, [][]any:
		DimensionPluckMap(target, data_arr, field.(string))
		return
	case []any:
		DimensionPluckStruct(target, data_arr, field.(string))
		return
	}
	DimensionPluckReflectSlice(target, data, field.(string))
}

// 二维数组某一个键，变一维数组
func DimensionPluckMap(target any, data any, field string) {
	if data == nil {
		return
	}
	switch target_arr := target.(type) {
	case *[]string:
		rangeMapAny(data, func(val any) {
			if val_val, ok := val.(string); ok {
				*target_arr = append(*target_arr, val_val)
			}
		}, field)
		return
	case *[]int:
		rangeMapAny(data, func(val any) {
			*target_arr = append(*target_arr, int(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int64:
		rangeMapAny(data, func(val any) {
			*target_arr = append(*target_arr, (convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int32:
		rangeMapAny(data, func(val any) {
			*target_arr = append(*target_arr, int32(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int16:
		rangeMapAny(data, func(val any) {
			*target_arr = append(*target_arr, int16(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int8:
		rangeMapAny(data, func(val any) {
			*target_arr = append(*target_arr, int8(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]any:
		rangeMapAny(data, func(val any) {
			*target_arr = append(*target_arr, (val))
		}, field)
		return
	}
}

func DimensionPluckStruct(target any, data []any, field string) {
	if data == nil || len(data) == 0 {
		return
	}
	switch target_arr := target.(type) {
	case *[]string:
		_rangeAny(data, func(val any) {
			if val_val, ok := val.(string); ok {
				*target_arr = append(*target_arr, (val_val))
			}
		}, field)
		return
	case *[]int:
		_rangeAny(data, func(val any) {
			*target_arr = append(*target_arr, int(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int64:
		_rangeAny(data, func(val any) {
			*target_arr = append(*target_arr, (convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int32:
		_rangeAny(data, func(val any) {
			*target_arr = append(*target_arr, int32(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int16:
		_rangeAny(data, func(val any) {
			*target_arr = append(*target_arr, int16(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int8:
		_rangeAny(data, func(val any) {
			*target_arr = append(*target_arr, int8(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]any:
		_rangeAny(data, func(val any) {
			*target_arr = append(*target_arr, (val))
		}, field)
		return
	}
}

func DimensionPluckReflectSlice(target any, data any, field string) {
	if data == nil {
		return
	}
	datat := reflect.TypeOf(data)
	valuet := reflect.ValueOf(data)
	if datat.Kind() == reflect.Ptr {
		datat = datat.Elem()
		valuet = valuet.Elem()
	}
	if datat.Kind() != reflect.Slice {
		return
	}
	switch target_arr := target.(type) {
	case *[]string:
		_rangeWithReflect(valuet, func(val any) {
			if val_val, ok := val.(string); ok {
				*target_arr = append(*target_arr, (val_val))
			}
		}, field)
		return
	case *[]int:
		_rangeWithReflect(valuet, func(val any) {
			*target_arr = append(*target_arr, int(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int64:
		_rangeWithReflect(valuet, func(val any) {
			*target_arr = append(*target_arr, (convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int32:
		_rangeWithReflect(valuet, func(val any) {
			*target_arr = append(*target_arr, int32(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int16:
		_rangeWithReflect(valuet, func(val any) {
			*target_arr = append(*target_arr, int16(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]int8:
		_rangeWithReflect(valuet, func(val any) {
			*target_arr = append(*target_arr, int8(convtype.Any2Int64(val)))
		}, field)
		return
	case *[]any:
		_rangeWithReflect(valuet, func(val any) {
			*target_arr = append(*target_arr, (val))
		}, field)
		return
	}
}

// internal functions

func rangeMapAny(data any, fn func(v any), field any) {
	switch datad := data.(type) {
	case []map[string]any:
		_rangeMapAnyString(datad, fn, field.(string))
	case []map[int]any:
		_rangeMapAnyInt(datad, fn, int(convtype.Any2Int64(field)))
	case [][]any:
		_rangeMapAnySlice(datad, fn, int(convtype.Any2Int64(field)))
	}
}
func _rangeMapAnyString(data []map[string]any, fn func(m any), field string) {
	for i, n := 0, len(data); i < n; i++ {
		item := data[i]
		val, ok := item[field]
		if !ok {
			continue
		}
		fn(val)
	}
}

func _rangeMapAnyInt(data []map[int]any, fn func(m any), field int) {
	for i, n := 0, len(data); i < n; i++ {
		item := data[i]
		val, ok := item[field]
		if !ok {
			continue
		}
		fn(val)
	}
}

func _rangeMapAnySlice(data [][]any, fn func(m any), field int) {
	for i, n := 0, len(data); i < n; i++ {
		item := data[i]
		if n > field {
			continue
		}
		fn(item)
	}
}

func _rangeAny(data []any, fn func(m any), field string) {
	for i, n := 0, len(data); i < n; i++ {
		item := reflect.TypeOf(data[i])
		value := reflect.ValueOf(data[i])
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
			value = value.Elem()
		}
		if item.Kind() != reflect.Struct {
			continue
		}
		field_value := value.FieldByName(field)
		if !field_value.IsValid() {
			continue
		}
		val := field_value.Interface()
		//if !ok {
		//	continue
		//}
		fn(val)
	}
}

func _rangeWithReflect(values reflect.Value, fn func(any), field string) {
	for i := 0; i < values.Len(); i++ {
		item := values.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}
		if item.Kind() != reflect.Struct {
			continue
		}
		v := item.FieldByName(field)
		if !v.IsValid() {
			continue
		}
		fn(v.Interface())
	}
}
