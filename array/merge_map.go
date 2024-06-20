package array

// 覆盖式 合并 map[string]stirng
func mergeMapString(data map[string]string, data2 map[string]string) map[string]string {
	if data2 == nil || len(data2) == 0 {
		return data
	}
	if data == nil {
		return data2
	}
	for k, v := range data2 {
		data[k] = v
	}
	return data
}

func MergeMapString(data map[string]string, data_arr ...map[string]string) map[string]string {
	if data_arr == nil || len(data_arr) == 0 {
		return data
	}
	for _, m := range data_arr {
		data = mergeMapString(data, m)
	}
	return data
}

func mergeMapStringAny(data map[string]any, data2 map[string]any) map[string]any {
	if data2 == nil || len(data2) == 0 {
		return data
	}
	if data == nil {
		return data2
	}
	for k, v := range data2 {
		data[k] = v
	}
	return data
}

func MergeMapStringAny(data map[string]any, data_arr ...map[string]any) map[string]any {
	if data_arr == nil || len(data_arr) == 0 {
		return data
	}
	for _, m := range data_arr {
		data = mergeMapStringAny(data, m)
	}
	return data
}
