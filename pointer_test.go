package tips

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIsNil(t *testing.T) {
	b := new(string)
	c := "a"
	b = nil
	var d = &Strsss{}
	if IsNil(b) == false {
		t.Error("b should be nil")
	}
	d = nil
	if IsNil(d) == false {
		t.Error("d should be nil")
	}
	if IsNil(c) == true {
		t.Error("c is not an empty string")
	}
}

func TestConfigWithViper(t *testing.T) {
	apath, _ := os.Getwd()

	vf := ConfigWithViper(apath + "/tmp/fix.yaml")
	lg := vf.Sub("Log")
	assert.True(t, lg.GetInt("Level") > 0)
}

type Strsss struct {
}

func Error() string {
	return ""
}
