package licenses

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestRasBock(t *testing.T) {
	max := big.NewInt(100)
	bi, err := rand.Int(rand.Reader, max)
	fmt.Println("rand.Reader", bi, err)

	rs := NewRsaKey(func(rsaa *RsaKeyBlock) {
		rsaa.GenerateByReader(rand.Reader)
	})
	if err != nil {
		printStack()
		t.Fatal("NewRsaKeyWithLength create RSA key error ", err)
	}
	rs.Path("./tmp")
	rs.WritePrivateKey()
	rs.WritePublicKey()
}

func TestLicense(t *testing.T) {
	//gob.Register(&License{})
	rs := NewRsaKey(
		OptionPrivateKeyFromFile("./tmp/private.key"),
		OptionPublicKeyFromFile("./tmp/public.pem"))
	rs.Path("./tmp/test")

	lserv := NewLicenseServ(func(serv *LicenseSrv) {
		serv.Model = MODEL_JSON
	}, func(serv *LicenseSrv) {
		serv.RSA = rs
		serv.Auth("admin", "testing")
		serv.Path("./tmp")
	})
	lserv.RSA.WritePublicKey() // 测试写入的RSA 公钥 是否 正确
	lobj := &License{
		AppId:      "admin",
		Secret:     "testing",
		Expiration: time.Now().Add(time.Hour),
		Timestamp:  time.Now().Unix(),
	}
	// make license
	lserv.MakeLicense(lobj)
	t.Log("MakeLicense OK ")

	lserv.RSA.DisallowPrivate()
	err := lserv.VerifyLicenseFileValid("./tmp/LICENSE")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Verify License OK")
}

// 测试license 被修改了 是否允许通过
func TestLicenseChange(t *testing.T) {
	rs := NewRsaKey(
		OptionPublicKeyFromFile("./tmp/public.pem"))
	rs.Path("./tmp/test")
	lserv := NewLicenseServ(func(serv *LicenseSrv) {
		serv.Model = MODEL_GOB
	}, func(serv *LicenseSrv) {
		serv.RSA = rs
		serv.Auth("admin", "testing")
		serv.Path("./tmp")
	})
	// 修改 license data  begin ------
	licenseData, err := os.ReadFile("./tmp/LICENSE")
	if err != nil {
		t.Fatal(errors.New("EN:(not found license file) CN:(未找到许可证文件)"))
	}
	var license License
	lserv.readByJson(&license, licenseData)
	license.Timestamp = time.Now().Add(time.Hour).Unix()
	licenseData, err = lserv.readJsonByte(&license)
	// 修改 license data  over -------

	err = lserv.VerifyLicenseValid(licenseData)
	if err == nil {
		t.Fatal(errors.New("the license of changed should not be passed!"))
	}
	t.Log("Verify License OK")
}

// 测试不同平台的license
func TestDiffPlatLicense(t *testing.T) {

}

// testing functions
func simpleSecret() string {

	b := make([]byte, 2048)
	//ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	fmt.Println(b)
	return string(b)
}

// 打印堆栈的函数
func printStack() {
	// 获取当前 goroutine 的调用栈信息
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // 跳过 printStack 和 runtime.Callers 本身
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		fmt.Printf("%s +%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}
