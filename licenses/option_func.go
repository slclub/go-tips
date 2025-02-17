package licenses

import (
	"encoding/base64"
	"errors"
	"github.com/slclub/go-tips/logf"
	"os"
	"time"
)

// functions
// license 基础检测
func OptionCheckLicenseBase(lrv *LicenseSrv, license *License) error {
	// 检查过期时间 这里是使用本地时间校验的，也可以通过网络时间校验
	if time.Now().After(license.Expiration) {
		return errors.New("EN:(license verification expired) CN:(许可证已过期) ")
	}

	// 检查用户名
	if !(license.AppId == lrv.license.AppId && license.Secret == lrv.license.Secret) {
		return errors.New("EN:(license verification unvalid) CN:(许可证无效) ")
	}
	return nil
}

// RSA ============================BEGIN============================
// 从文件中 读取 RSA的 公钥
func OptionPublicKeyFromFile(filename string) func(serv *RsaKeyBlock) {
	return func(serv *RsaKeyBlock) {
		data, err := os.ReadFile(filename)
		if err != nil {
			logf.Log().Printf("[ERROR] read file:%v error:%v", filename, err)
			return
		}
		err = serv.PublicKeyRead(data)
		//logf.Log().Printf("[Info]license.PublicKeyRead OK")
	}
}

// 从文件中 读取 RSA的 私钥
func OptionPrivateKeyFromFile(filename string) func(serv *RsaKeyBlock) {
	return func(serv *RsaKeyBlock) {
		data, err := os.ReadFile(filename)
		if err != nil {
			logf.Log().Printf("[ERROR] read file:%v error:%v", filename, err)
			return
		}
		err = serv.PrivateKeyRead(data)
		//logf.Log().Printf("[Info]license.PrivateKeyRead OK")
	}
}

// RSA ============================ END ============================

// common functions
func Base64Encode(data string) (string, error) {
	// 将字符串转换为[]byte
	dataBytes := []byte(data)
	// 编码
	encoded := base64.StdEncoding.EncodeToString(dataBytes)
	return encoded, nil
}

func Base64Decode(encoded string) (string, error) {
	// 解码
	dataBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	// 将[]byte转换回字符串
	data := string(dataBytes)
	return data, nil
}
