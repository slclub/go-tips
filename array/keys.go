package array

func KeysMapStringAny(data map[string]any) []string {
	if data == nil {
		return nil
	}
	rtn := []string{}
	for k, _ := range data {
		rtn = append(rtn, k)
	}
	return rtn
}

func MapKeys(data any) []string {
	if data == nil {
		return nil
	}
	rtn := []string{}
	switch arr := data.(type) {
	case map[string]any:
		return KeysMapStringAny(arr)
	case map[string]int:
		for k, _ := range arr {
			rtn = append(rtn, k)
		}
	case map[string]string:
		for k, _ := range arr {
			rtn = append(rtn, k)
		}
	}
	if len(rtn) == 0 {
		return nil
	}
	return rtn
}
