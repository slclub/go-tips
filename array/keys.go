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
