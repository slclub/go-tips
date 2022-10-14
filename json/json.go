package json

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func Decode(src, target any) error {
	var read io.Reader
	switch val := src.(type) {
	case string:
		read = strings.NewReader(val)
	case []byte:
		read = bytes.NewBuffer(val)
	}
	decoder := json.NewDecoder(read)
	decoder.UseNumber()
	err := decoder.Decode(&target)
	return err
}

func Unmarshal(b []byte, target any) error {
	return json.Unmarshal(b, target)
}

func Marshal(val any) ([]byte, error) {
	return json.Marshal(val)
}
