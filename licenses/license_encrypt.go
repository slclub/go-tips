// 使用私钥加密 生成license
package licenses

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

const (
	MODEL_GOB  = "gob"
	MODEL_JSON = "json"
)

type License struct {
	AppId      string    `json:"app_id"`              // 用户名
	Secret     string    `json:"secret"`              // 密钥
	Expiration time.Time `json:"expiration"`          // 到期时间
	Timestamp  int64     `json:"timestamp"`           // 系统时间戳
	Signature  []byte    `json:"signature,omitempty"` // 签名
}

func New(rs *RsaKeyBlock) *LicenseSrv {
	obj := &LicenseSrv{RSA: rs}
	return obj
}

type LicenseSrv struct {
	license_path string
	RSA          *RsaKeyBlock
	license      *License
	Model        string
	CheckHandle  []func(srv *LicenseSrv, license *License) error
}

//
//var (
//	license = &License{
//		Username:   Username,
//		Secret:     Secret,
//		Expiration: time.Now().Add(30 * 24 * time.Hour), // 30 天后过期
//	}
//)

func NewLicenseServ(fns ...func(*LicenseSrv)) *LicenseSrv {
	serv := &LicenseSrv{
		Model: MODEL_JSON,
	}
	for _, fn := range fns {
		fn(serv)
	}
	serv.CheckHandle = append(serv.CheckHandle, OptionCheckLicenseBase)
	return serv
}

func (self *LicenseSrv) Auth(appid, secret string) {
	if self.license == nil {
		self.license = &License{}
	}
	self.license.AppId = appid
	self.license.Secret = secret
}

func (self *LicenseSrv) MakeLicense(license *License) error {
	// 签名许可证
	signature, err := self.signLicense(license, self.RSA.PrivateKey())
	if err != nil {
		return err
	}
	license.Signature = signature

	return self.save(self.Path()+"/"+"LICENSE", license)
}

func (self *LicenseSrv) MakeLicenseSimple(license *License) error {
	// 签名许可证
	signature, err := self.signLicense(license, self.RSA.PrivateKey())
	if err != nil {
		return err
	}
	license.Signature = signature

	return nil
}

func (self *LicenseSrv) Path(ps ...string) string {
	if len(ps) > 0 {
		self.license_path = ps[0]
	}
	if n := len(self.license_path); n > 0 && self.license_path[n-1] == '/' {
		self.license_path = self.license_path[:n-1]
	}
	return self.license_path
}

func (self *LicenseSrv) Conv2Bytes(data *License) ([]byte, error) {
	// 方式2 加入缓冲区
	var buf1 bytes.Buffer

	switch self.Model {
	case MODEL_JSON:
		enc := json.NewEncoder(&buf1)
		err := enc.Encode(data)
		return buf1.Bytes(), err
	default:
		enc := gob.NewEncoder(&buf1)
		err := enc.Encode(data)
		return buf1.Bytes(), err
	}
}

// protected methods .
// 签名许可证
func (self *LicenseSrv) signLicense(license *License, privateKey *rsa.PrivateKey) ([]byte, error) {
	licenseCopy := *license
	licenseCopy.Signature = nil // 签名前移除签名字段

	// 使用gob编码
	licenseData, err := self.readGobByte(&licenseCopy)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(licenseData)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func (self *LicenseSrv) save(filename string, data any) error {
	switch self.Model {
	case MODEL_JSON:
		return self.saveJson(filename, data)
	default:
		return self.saveGob(filename, data)
	}
}

func (self *LicenseSrv) saveGob(filename string, data any) error {
	err := os.Remove(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	//方式1
	//fs, err := os.Create(filename)
	//if err != nil {
	//	return err
	//}
	//defer fs.Close()
	//enc := gob.NewEncoder(fs)
	//if err := enc.Encode(data); err != nil {
	//	panic(err)
	//}

	// 方式2 加入缓冲区
	var buf1 bytes.Buffer
	enc := gob.NewEncoder(&buf1)
	err = enc.Encode(data)
	if err != nil {
		panic(err)
	}
	// write to file
	err = ioutil.WriteFile(filename, buf1.Bytes(), 0600)
	if err != nil {
		panic(err)
	}

	return nil
}

func (self *LicenseSrv) saveJson(filename string, data any) error {
	err := os.Remove(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	var jdata bytes.Buffer
	enc := json.NewEncoder(&jdata)
	err = enc.Encode(data)
	if err != nil {
		panic(err)
	}
	// write to file
	err = ioutil.WriteFile(filename, jdata.Bytes(), 0600)
	if err != nil {
		panic(err)
	}

	return nil
}
