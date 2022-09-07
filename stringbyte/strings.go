package stringbyte

// get left string of s
func SubLeft(s, sep string) string {
	ls := len(sep)
	if len(s) == 0 || ls == 0 {
		return s
	}

	for i, _ := range s {
		j := i + 1
		if j < ls {
			continue
		}
		if s[j-ls:j] != sep {
			continue
		}
		return s[:j-ls]
	}
	return s
}
