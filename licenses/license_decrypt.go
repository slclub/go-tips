package licenses

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// 验证许可证签名
func (self *LicenseSrv) VerifyLicense(license *License, publicKey *rsa.PublicKey) error {
	licenseCopy := *license
	licenseCopy.Signature = nil

	licenseData, err := self.readGobByte(&licenseCopy)
	if err != nil {
		return err
	}
	if publicKey == nil {
		return errors.New("RSA Verify public key is nil")
	}
	hash := sha256.Sum256(licenseData)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], license.Signature)
}

func (self *LicenseSrv) VerifyLicenseFileValid(filename string) error {
	licenseData, err := os.ReadFile(filename)
	if err != nil {
		return errors.New("EN:(not found license file) CN:(未找到许可证文件)")
	}
	return self.VerifyLicenseValid(licenseData)
}

func (self *LicenseSrv) VerifyLicenseValid(licenseData []byte) error {

	var license License
	err := self.ReadBy(&license, licenseData)
	if err != nil {
		//fmt.Println("解析许可证文件失败")
		return err
	}
	//fmt.Println("VerifyLicenseValid read ", license)
	// 验证签名
	err = self.VerifyLicense(&license, self.RSA.PublicKey())
	if err != nil {
		return errors.New("EN:(license verification failed) CN:(许可证验证失败) " + err.Error())
	}

	// 自定义检查
	for _, fn := range self.CheckHandle {
		err = fn(self, &license)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *LicenseSrv) GetLicenseFromInternet(license *License, apiUrl string, appId string, fn func(bd []byte) []byte) error {
	resp, err := http.Post(apiUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(`{"AppId": "`+appId+`"}`))
	if err != nil {
		log.Fatal("httpPost  error:" + err.Error())
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatal("no resp  error:" + err.Error())
		return err
	}
	body = fn(body)
	return self.ReadBy(license, body)
}

func (self *LicenseSrv) ReadBy(license *License, data []byte) error {
	switch self.Model {
	case MODEL_JSON:
		return self.readByJson(license, data)
	default:
		return self.readByGob(license, data)
	}
}

func (self *LicenseSrv) readByGob(license *License, data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&license)
	if err != nil {
		//fmt.Println("解析许可证文件失败")
		return errors.New("EN:(failed to parse license) CN:(解析许可证文件失败)")
	}
	return nil
}

func (self *LicenseSrv) readByJson(license *License, data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)
	err := decoder.Decode(&license)
	if err != nil {
		//fmt.Println("解析许可证文件失败")
		return errors.New("EN:(failed to parse license) CN:(解析许可证文件失败)")
	}
	return nil
}

func (self *LicenseSrv) ReadConvToByte(license *License) ([]byte, error) {
	switch self.Model {
	case MODEL_JSON:
		return self.readJsonByte(license)
	default:
		return self.readGobByte(license)
	}
}

func (self *LicenseSrv) readGobByte(license *License) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(*license)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (self *LicenseSrv) readJsonByte(license *License) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(*license)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
