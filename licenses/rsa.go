package licenses

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

const (
	CONST_TYPE_PRIVATE = "RSA PRIVATE KEY"
	CONST_TYPE_PUBLIC  = "RSA PUBLIC KEY"
	DEFAULT_KEN_LEN    = 2048
)

type RsaKeyBlock struct {
	pub_key             *rsa.PublicKey
	pri_key             *rsa.PrivateKey
	key_path            string
	can_use_private_key bool
	Bits                int
}

func NewRsaKeyWithLength(mk string, lenghs ...int) (*RsaKeyBlock, error) {
	l := DEFAULT_KEN_LEN
	if len(lenghs) > 0 {
		l = lenghs[0]
	}
	obj := &RsaKeyBlock{
		can_use_private_key: true,
		Bits:                l,
	}
	err := obj.GenerateByString(mk, lenghs...)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func NewRsaKey(fns ...func(*RsaKeyBlock)) *RsaKeyBlock {
	obj := &RsaKeyBlock{
		can_use_private_key: true,
		Bits:                DEFAULT_KEN_LEN,
	}
	for _, fn := range fns {
		fn(obj)
	}
	return obj
}

// 从io.Reader 生成 私钥 - 公钥
func (self *RsaKeyBlock) GenerateByReader(reader io.Reader) error {

	// 生成 RSA 密钥对
	//a := rand.Reader
	privateKey, err := rsa.GenerateKey(reader, self.Bits)
	if err != nil {
		return err
	}

	self.pri_key = privateKey // 生成时，强制允许赋值 因此不用 self.PrivateKey(xxx)

	self.PublicKey(&self.pri_key.PublicKey)
	return nil
}

func (self *RsaKeyBlock) GenerateByString(mk string, lenghs ...int) error {
	buffer := bytes.Buffer{}
	buffer.WriteString(mk)
	//err := binary.Write(&buffer, binary.BigEndian, uint32(mk))
	//if err != nil {
	//	logf.Log().Printf("Error encoding number:", err)
	//	return err
	//}
	return self.GenerateByReader(&buffer)
}

func (self *RsaKeyBlock) PublicKey(kkk ...*rsa.PublicKey) *rsa.PublicKey {
	if len(kkk) > 0 {
		self.pub_key = kkk[0]
	}
	return self.pub_key
}

func (self *RsaKeyBlock) PrivateKey(kkk ...*rsa.PrivateKey) *rsa.PrivateKey {
	if !self.can_use_private_key {
		return nil
	}
	if len(kkk) > 0 {
		self.pri_key = kkk[0]
	}
	return self.pri_key
}

func (self *RsaKeyBlock) DisallowPrivate() {
	self.can_use_private_key = false
}

func (self *RsaKeyBlock) Key2Bytes(arg any) []byte {
	switch val := arg.(type) {
	case *rsa.PrivateKey:
		return self.privateKey2Bytes(val)
	case *rsa.PublicKey:
		return self.publicKey2Bytes(val)
	}
	return nil
}

func (self *RsaKeyBlock) Path(ps ...string) string {
	if len(ps) > 0 {
		self.key_path = ps[0]
	}
	if n := len(self.key_path); n > 0 && self.key_path[n-1] == '/' {
		self.key_path = self.key_path[:n-1]
	}
	return self.key_path
}

func (self *RsaKeyBlock) WriteAll() error {
	err := self.WritePrivateKey()
	if err != nil {
		return err
	}
	err = self.WritePublicKey()
	return err
}

func (self *RsaKeyBlock) WritePrivateKey() error {
	filename := self.Path() + "/" + "private.key"
	block := &pem.Block{
		Type:  CONST_TYPE_PRIVATE,
		Bytes: self.Key2Bytes(self.PrivateKey()),
	}
	err := self.save(filename, block)
	if err == nil {
		return nil
	}
	return errors.New(block.Type + ":" + err.Error())
}

func (self *RsaKeyBlock) WritePublicKey() error {
	filename := self.Path() + "/" + "public.pem"
	block := &pem.Block{
		Type:  CONST_TYPE_PUBLIC,
		Bytes: self.Key2Bytes(self.PublicKey()),
	}
	err := self.save(filename, block)
	if err == nil {
		return nil
	}
	return errors.New(block.Type + ":" + err.Error())
}

func (self *RsaKeyBlock) PrivateKeyRead(p []byte) error {
	privateBlock, _ := pem.Decode(p)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		return (err)
	}
	self.pri_key = privateKey
	return nil
}

func (self *RsaKeyBlock) PublicKeyRead(p []byte) error {
	publicBlock, _ := pem.Decode(p)
	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		return err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("EN(Not RSA Public Key)  CN(不是 RSA 公钥)")
	}
	self.PublicKey(rsaPublicKey)
	return nil
}

// protected methods
func (self *RsaKeyBlock) privateKey2Bytes(key *rsa.PrivateKey) []byte {
	if key == nil {
		panic(errors.New("private key is nil"))
	}
	return x509.MarshalPKCS1PrivateKey(key)
}

func (self *RsaKeyBlock) publicKey2Bytes(key *rsa.PublicKey) []byte {
	if key == nil {
		panic(errors.New("public key is nil"))
	}
	ret, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		panic(err)
	}
	return ret
}

func (self *RsaKeyBlock) save(filename string, block *pem.Block) error {
	err := os.Remove(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	fs, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fs.Close()
	if err := pem.Encode(fs, block); err != nil {
		return err
	}
	return nil
}

// config functions
