package date

import (
	"fmt"
	"testing"
)

func TestStrToTimestamp(t *testing.T) {
	str := "2022-09-30 11:20"
	t1 := StrToTimestamp(str)
	fmt.Println(t1)
	fmt.Println(TimestampToStr(t1))
}
