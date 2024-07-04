package convtype

import (
	"github.com/slclub/go-tips/stringbyte"
	"io"
	"strconv"
)

type StringPrinter interface {
	String() string
}

func String(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case int64:
		return strconv.FormatInt(val, 10)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case uint64:
		return strconv.FormatInt(int64(val), 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case uint32:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case uint16:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint8:
		return strconv.FormatInt(int64(val), 10)
	case []byte:
		return string(val)
	case []rune:
		return string(val)
	case io.Reader:
		data, err := io.ReadAll(val)
		if err == nil {
			return stringbyte.BytesToString(data)
		}
	case StringPrinter:
		return val.String()
	case error:
		return val.Error()
	}
	return ""
}
