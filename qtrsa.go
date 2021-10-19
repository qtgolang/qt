package qt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type Qtrsa struct {
	privateKey *pem.Block
	publicKey  *pem.Block
}

//Encrypt
// 加密
func (c *Qtrsa) Encrypt(origData []byte) ([]byte, error) {
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(c.publicKey.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//SetPrivateKey
// 设置私钥
func (c *Qtrsa) SetPrivateKey(privateKey string) error {
	c.privateKey, _ = pem.Decode([]byte(privateKey))
	if c.privateKey == nil {
		return errors.New("public key error")
	}
	return nil
}

//SetPublicKey
// 设置公钥
func (c *Qtrsa) SetPublicKey(publicKey string) error {
	c.publicKey, _ = pem.Decode([]byte(publicKey))
	if c.publicKey == nil {
		return errors.New("public key error")
	}
	return nil
}

//Decrypt
// 解密
func (c *Qtrsa) Decrypt(ciphertext []byte) ([]byte, error) {
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(c.privateKey.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//Tnit
// 设置 默认的公钥私钥
func (c *Qtrsa) Tnit() {
	//可通过openssl产生
	//openssl genrsa -out rsa_private_key.pem 1024
	//openssl
	//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
	c.SetPublicKey(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB`)
	c.SetPublicKey("\r\n-----BEGIN RSA PRIVATE KEY-----\r\nMIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQABAoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg10F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiFu+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYyXhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5M0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0SpxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4py2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW\r\n-----END RSA PRIVATE KEY-----\r\n")
}

