package cert

import (
	"github.com/qtgolang/qt/amiddleman/GoIyov/cache"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	rootCa  *x509.Certificate // CA证书
	rootKey *rsa.PrivateKey   // 证书私钥
)

var certCache *cache.Cache
var RootKey []byte
var RootCa []byte

func Init(rootCa, rootKey string) {
	RootCa = []byte(rootCa)
	RootKey = []byte(rootKey)
	certCache = cache.NewCache()
	if err := loadRootCa(); err != nil {
		panic(err)
	}
	if err := loadRootKey(); err != nil {
		panic(err)
	}
}

func GetCertificate(host string) (tls.Certificate, error) {
	certificate, err := certCache.GetOrStore(host, func() (interface{}, error) {
		host, _, err := net.SplitHostPort(host)
		if err != nil {
			return nil, err
		}
		certByte, priByte, err := generatePem(host)
		if err != nil {
			return nil, err
		}
		certificate, err := tls.X509KeyPair(certByte, priByte)
		if err != nil {
			return nil, err
		}
		return certificate, nil
	})
	return certificate.(tls.Certificate), err
}
func generatePem(host string) ([]byte, []byte, error) {
	max := new(big.Int).Lsh(big.NewInt(1), 128)   //把 1 左移 128 位，返回给 big.Int
	serialNumber, _ := rand.Int(rand.Reader, max) //返回在 [0, max) 区间均匀随机分布的一个随机值
	template := x509.Certificate{
		SerialNumber: serialNumber, // SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		Subject: pkix.Name{ //Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略。
			CommonName: host,
		},
		NotBefore:      time.Now().AddDate(-1, 0, 0),
		NotAfter:       time.Now().AddDate(1, 0, 0),
		KeyUsage:       x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, //KeyUsage 与 ExtKeyUsage 用来表明该证书是用来做服务器认证的
		ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // 密钥扩展用途的序列
		EmailAddresses: []string{"forward.nice.cp@gmail.com"},
	}

	if ip := net.ParseIP(host); ip != nil {
		template.IPAddresses = []net.IP{ip}
	} else {
		template.DNSNames = []string{host}
	}

	priKey, err := generateKeyPair()
	if err != nil {
		return nil, nil, err
	}

	cer, err := x509.CreateCertificate(rand.Reader, &template, rootCa, &priKey.PublicKey, rootKey)
	if err != nil {
		return nil, nil, err
	}

	return pem.EncodeToMemory(&pem.Block{ // 证书
			Type:  "CERTIFICATE",
			Bytes: cer,
		}), pem.EncodeToMemory(&pem.Block{ // 私钥
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priKey),
		}), err
}

// 秘钥对 生成一对具有指定字位数的RSA密钥
func generateKeyPair() (*rsa.PrivateKey, error) {
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.Wrap(err, "密钥对生成失败")
	}

	return priKey, nil
}

// 加载根证书
func loadRootCa() error {
	p, _ := pem.Decode(RootCa)
	var err error
	rootCa, err = x509.ParseCertificate(p.Bytes)
	if err != nil {
		return errors.Wrap(err, "CA证书解析失败")
	}

	return nil
}

// 加载根Private Key
func loadRootKey() error {
	p, _ := pem.Decode(RootKey)
	var err error
	rootKey, err = x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return errors.Wrap(err, "Key证书解析失败")
	}

	return err
}

// 获取证书原内容
func GetCaCert() []byte {
	return RootCa
}

// 添加信任跟证书至钥匙串
func AddTrustedCert() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := dir + "/caRoot.crt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer os.Remove(fileName)
	defer file.Close()

	file.Write(RootCa)

	var command string
	switch runtime.GOOS {
	case "darwin":
		command = fmt.Sprintf("sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain %s", fileName)
	case "windows":
		command = fmt.Sprintf("certutil -addstore -f \"ROOT\" %s", dir+"\\caRoot.crt")
	default:
		return errors.New("仅支持MaxOS/Windows系统")
	}

	return shell(command)
}

// 执行shell命令
func shell(command string) error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "-c", command)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			return errors.Wrap(err, "")
		}
		return errors.Wrap(cmd.Wait(), out.String())
	}
	cmd := exec.Command("sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return errors.Wrap(err, "")
	}
	return errors.Wrap(cmd.Wait(), out.String())
}
