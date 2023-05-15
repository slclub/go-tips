package convtype

import (
	"github.com/slclub/go-tips/logf"
	"github.com/slclub/go-tips/stringbyte"
	"strconv"
)

func Any2Int64(v any) int64 {
	switch val := v.(type) {
	case int8:
		return int64(val)
	case uint8:
		return int64(val)
	case int32:
		return int64(val)
	case uint32:
		return int64(val)
	case uint:
		return int64(val)
	case int16:
		return int64(val)
	case uint16:
		return int64(val)
	case uint64:
		return int64(val)
	case int64:
		return val
	case float64:
		return int64(val)
	case float32:
		return int64(val)
	case int:
		return int64(val)
	case []byte:
		n, err := strconv.ParseInt(stringbyte.BytesToString(val), 10, 64)
		if err != nil {
			logf.Print("TIPS.WARN Any2Int64 err:", err)
		}
		return n
	case string:
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			logf.Print("TIPS.WARN Any2Int64 err:", err)
		}
		return n
	}
	return 0
}
