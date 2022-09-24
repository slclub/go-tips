package date

import "time"

const (
	FORMAT_DATE = "2006-01-02 15:04:05"
)

func StrToTimestamp(str string) int64 {
	df := "1970-01-01 00:00:00"
	pl := len(str)
	dl := len(df)
	if pl < dl {
		str += df[pl:]
	}
	t, err := time.ParseInLocation(FORMAT_DATE, str, time.Local)
	if err == nil {
		return t.Unix()
	}
	return 0
}

func TimestampToStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(FORMAT_DATE)
}

func Unix(args ...int64) time.Time {
	switch len(args) {
	case 1:
		return time.Unix(args[0], 0)
	case 2:
		return time.Unix(args[0], args[1])
	}
	return time.Unix(0, 0)
}
