package tips

import "github.com/slclub/go-tips/convtype"

// convert to string
func String(v any) string {
	return convtype.String(v)
}

// get sub string index of string
func StrPos(str, sep string) int {
	sepl := len(sep)
	if len(str) < sepl {
		return -1
	}
	for i, n := 0, len(str); i < n; i++ {
		if i+3 >= n {
			return -1
		}
		if i+sepl > n {
			return -1
		}
		if str[i:i+sepl] == sep {
			return i
		}
	}
	return -1
}

// get prefix string that the sub string index of string
func StrBegin(str string, sep string) string {
	if sep == "" {
		return str
	}
	i := StrPos(str, sep)
	if i == -1 {
		return str
	}
	return str[:i]
}

// get suffix string that the sub string index of string
func StrEnd(str, sep string) string {
	if sep == "" {
		return ""
	}
	sepl := len(sep)
	i := StrPos(str, sep)
	if i == -1 {
		return ""
	}
	return str[i+sepl:]
}
