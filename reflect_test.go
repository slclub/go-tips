package tips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _domain = "github.com/slclub/go-tips"

func Test_FUNC_NAME(t *testing.T) {
	fn := FUNC_NAME(needGetMyNameFunc)
	assert.Equal(t, _domain+".needGetMyNameFunc", fn)
}

func Test_IS_FUNC(t *testing.T) {
	fn := needGetMyNameFunc
	assert.True(t, IS_FUNC(fn))
	fn2 := "string"
	assert.False(t, IS_FUNC(fn2))
}

func Test_IS_STRUCT(t *testing.T) {
	u1 := stu{}
	assert.True(t, IS_STRUCT(u1))
	assert.False(t, IS_STRUCT(2))
	assert.False(t, IS_STRUCT("I am not "))
}

func Test_IS_SLICE(t *testing.T) {
	s1 := []int{1, 2, 3, 2}
	assert.True(t, IS_SLICE(s1))
	assert.False(t, IS_SLICE(2))
}

func needGetMyNameFunc() {}

type stu struct{}
