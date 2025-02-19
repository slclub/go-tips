package licenses

import (
	"errors"
	"github.com/slclub/go-tips/logf"
	"io"
	"strconv"
	"strings"
	"time"
)

const CONST_SEP = "&"
const TIME_LAYOUT = "2006-01-02 15:04:05"

// 证书 接口
type Licenser interface {
	GetAppId() string
	GetSecret() string
	GetExpiration() time.Time
	String() string
	Bytes() []byte
	GetSignature() []byte
	SetSignature(sign []byte)
	Len() int
	io.Writer
	io.Reader
}

type License struct {
	AppId      string    // 用户名
	Secret     string    // 密钥
	Expiration time.Time // 到期时间
	Timestamp  int64     // 系统时间戳
	Signature  []byte    `json:"Signature,omitempty"` // 签名
	Extra      string
	Offset     int
}

func (self *License) GetAppId() string {
	return self.AppId
}

func (self *License) GetSecret() string {
	return self.Secret
}

func (self *License) GetExpiration() time.Time {
	return self.Expiration
}

func (self *License) GetSignature() []byte {
	return self.Signature
}

func (self *License) SetSignature(sign []byte) {
	self.Signature = sign
}

func (self *License) Write(p []byte) (n int, err error) {
	n = len(p)
	br := string(p)
	arr := strings.Split(br, CONST_SEP)
	data := map[string]any{}
	for _, sr := range arr {
		arrrs := strings.Split(sr, "=")
		if len(arrrs) < 2 {
			return 0, errors.New("Conv bytes to Licenser error!")
		}
		data[arrrs[0]] = arrrs[1]
	}
	self.convFromMap(data)
	return n, err
}

func (self *License) Read(p []byte) (int, error) {
	str := self.String()
	strByte := []byte(str)
	n := copy(p, strByte)
	self.Offset = n
	return n, nil
}

func (self *License) Len() int {
	return len(self.String())
}

func (self *License) Bytes() []byte {
	return []byte(self.String())
}

func (self *License) String() string {
	str := "AppId=" + self.AppId
	str += "Secret=" + self.Secret
	str += "Expiration=" + self.Expiration.Format(TIME_LAYOUT)
	str += "Timestamp=" + strconv.FormatInt(self.Timestamp, 10)
	str += "Extra=" + self.Extra
	return str
}

func (self *License) convFromMap(m map[string]any) {
	if val, ok := m["AppId"]; ok {
		self.AppId, _ = val.(string)
	}
	if val, ok := m["Secret"]; ok {
		self.Secret, _ = val.(string)
	}
	if val, ok := m["Secret"]; ok {
		self.Secret, _ = val.(string)
	}
	if val, ok := m["Extra"]; ok {
		self.Extra, _ = val.(string)
	}
	if val, ok := m[""]; ok {
		self.Expiration, _ = time.Parse(TIME_LAYOUT, val.(string))
	}
	if val, ok := m["Timestamp"]; ok {
		switch v := val.(type) {
		case int64:
			self.Timestamp = v
		case int:
		case uint64:
		case float64:
		case uint:
			self.Timestamp = int64(v)
		case string:
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				logf.Print("TIPS.WARN Any2Int64 err:", err)
			}
			self.Timestamp = n
		}
	}
}
