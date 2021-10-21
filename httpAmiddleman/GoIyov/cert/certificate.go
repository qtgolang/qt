package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/nicecp/GoIyov/cache"
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

var (
	_rootCa = []byte(`-----BEGIN CERTIFICATE-----
MIIDiTCCAnGgAwIBAgIUE/FkeKfrKhOspQrqC6eu5hAx0jwwDQYJKoZIhvcNAQEL
BQAwUzELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAkJKMQswCQYDVQQHDAJCSjEMMAoG
A1UECgwDcGxlMQ0wCwYDVQQLDARsaXZlMQ0wCwYDVQQDDARSb290MCAXDTIxMTAy
MDA5MzQwNloYDzIxMjEwOTI2MDkzNDA2WjBTMQswCQYDVQQGEwJDTjELMAkGA1UE
CAwCQkoxCzAJBgNVBAcMAkJKMQwwCgYDVQQKDANwbGUxDTALBgNVBAsMBGxpdmUx
DTALBgNVBAMMBFJvb3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQD3
PRGH6mwE7G+txkcM1OLN+o/RcY4syZtdPjhN1Kw26u4XIOttv9Bm26unh9PpTUWf
SeD0+U5/Dd8xFignLKbzcs69gNT64wHQfJSX5Fu+jgldXAvjes9hQs4jKnDDdXpb
FN779bKIddjBVKVEvXaSwXgZIq+kqP+QecBTXhFqEbHhqyzbvJE0at+u9NPaDX23
nZF9jwRNuuuufBf7dcCqjGHBUbeXZNmePeW6Qm3T4+zMP6kSNqV7q2exPA82zWAe
+cQLtls417ksyrRhopCzA1tuwk9BRr5S3AlNLrX3dlVYXzRK70KDYcGb3uYxPcWq
x7DfnxG3czyF4RClEkr/AgMBAAGjUzBRMB0GA1UdDgQWBBSCsoTDDBvKhKnYRDCM
mVDpOWONjDAfBgNVHSMEGDAWgBSCsoTDDBvKhKnYRDCMmVDpOWONjDAPBgNVHRMB
Af8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQDVaKjuAbwdcQEZ262TJuYKSvcY
uPWQPWq8zVp+EZxJVGiPJwuISdjK9rE8QZGdeuoeilBvmbL6k2ETq4K+y2GmEEpC
2dsj9nwflRDxFVQsRBA/6JAWA0+YUY3MRWKdfTux5f7m5RBGiGCtJAbBuh/ZitA9
xujXOa3qn1XZfqVPUdSQiV/A43/jo1rDCn0j539DILO7brIpRslpWSeGymt7ODQC
GnNoekL1c7h9uX9g3xkdbhkgdhh6eG58suzRsVoCsbCj3w/3QRdoUefzsw2cpkpF
wd/miJevRR8ncgc7RwUavkS/boGPWUixso88d5RSeqbAv6BR7VG13EO/Ydoz
-----END CERTIFICATE-----
`)
	_rootKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA9z0Rh+psBOxvrcZHDNTizfqP0XGOLMmbXT44TdSsNuruFyDr
bb/QZturp4fT6U1Fn0ng9PlOfw3fMRYoJyym83LOvYDU+uMB0HyUl+Rbvo4JXVwL
43rPYULOIypww3V6WxTe+/WyiHXYwVSlRL12ksF4GSKvpKj/kHnAU14RahGx4ass
27yRNGrfrvTT2g19t52RfY8ETbrrrnwX+3XAqoxhwVG3l2TZnj3lukJt0+PszD+p
Ejale6tnsTwPNs1gHvnEC7ZbONe5LMq0YaKQswNbbsJPQUa+UtwJTS6193ZVWF80
Su9Cg2HBm97mMT3Fqsew358Rt3M8heEQpRJK/wIDAQABAoIBAQCzKPLCPjLGopsP
HyaverlcMBz11kcD15iZShQ8+kdNiJK9+eIA5sXbM4ZBYaFDZ/ZyxnOYseybD29U
P80bDjVxJxn/oxMzNztCXHTFWPrOrFjG6YPH9V/ACEwemYubaE8hH4+yn8ofLt7C
wlb86Bq7oC0qccM8HCcOB9xBzWHrLLjN+Y5zJWFaJPUVY1vsNmZMorAOHzygLG8F
QJNaX1GGRXZSFtmPTvKEJSxigEwj3RRPVDePFHP4biViGXgNHR2tHB64Um4qef4z
ZbMdLbfodLAUTwL/D8Fc9Q3g4BuA/VMLPNV0XOmzHoYUHQjuxvfxtLeMPjCWnE/s
+TlrQ+q5AoGBAPyx+jzG/ccPcpW3h8Qd2Jaj7K73rqN45ZqyX8TSBDFX8Qd/Oooz
aA3+K6FN8/3LWnvz5QwXvhQ1xeph/riMc6ssA9jQrqL1C9TkgXZHShqJivuiKzu0
i36k4R38QZMnl7UiQ7MNZN0e5TgZ8k0i2KufCK7aPDyVdvZOW9RBhhNjAoGBAPp4
0nK/jKZi7TMPeHt0rEyiNzkZyOHnkf+xRAtNLdALDbFLpwDEZ499COlzDpgz06ab
lCQjWNa4Kz5WISPZ+ty5bG9p9uH0lnD6QlyPU1nVgTCvr5GynJ4fZNfjn5Hygodv
ATjyiQgsrW7pq0fviO2wcvqevz+D54HYUW+/MfK1AoGBAJZYDdo7SqI0vqf1GgHF
ACggP6GaG32HYJQ5rGEd0wDIoc8kE2BGVZJ9ttex2YkWhC9bXNtlBOJhW++nfjWu
2uLsvR0yi7TIttFjYuNMZvqC+v3b7n0HXjdrQcTlYN58n/ZU/JJ7VZd52kcWqOLb
6K2zYScnEM+63ZyN4nTWxz6hAoGAFODzcftDpy8B5Mq6WVgtcKno/oqGs0YRZoYJ
TQPe+MOjHY9X2XmFxHFAx+z+X3Oahf3cCHMl2ag6epTFaG6oObP/NP5ZRRaVX8+M
rpiH8yoX/c33Tabc5VVqm5Bu4cScWtvG909IWvUWc/NogrOV73JQ81E+UfYV2z4D
89O1Py0CgYEA5B9ec0WAQGZMSATwzMZfjjP9xPxwmaNpFhJEW2E3IlSfeyFxhh0S
jmsJX9MzML6wc7PcCvN5xLkPaHz7l1Vgz8x/gg264voQXwUp96iEl2U5QI7FYcK4
G5wZBLTO7tMXhsqOOLou2Px+k09iYAxgeF1DlNqpAy2JulkKhidT9DE=
-----END RSA PRIVATE KEY-----

`)
)

var certCache *cache.Cache

func init() {
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
	p, _ := pem.Decode(_rootCa)
	var err error
	rootCa, err = x509.ParseCertificate(p.Bytes)
	if err != nil {
		return errors.Wrap(err, "CA证书解析失败")
	}

	return nil
}

// 加载根Private Key
func loadRootKey() error {
	p, _ := pem.Decode(_rootKey)
	var err error
	rootKey, err = x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return errors.Wrap(err, "Key证书解析失败")
	}

	return err
}

// 获取证书原内容
func GetCaCert() []byte {
	return _rootCa
}

// 添加信任跟证书至钥匙串
func AddTrustedCert() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := dir + "/caRootCert.crt"
	fmt.Println(fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer os.Remove(fileName)
	defer file.Close()

	file.Write(_rootCa)

	var command string
	switch runtime.GOOS {
	case "darwin":
		command = fmt.Sprintf("sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain %s", fileName)
	case "windows":
		command = fmt.Sprintf("certutil -addstore -f \"ROOT\" %s", fileName)
	default:
		return errors.New("仅支持MaxOS/Windows系统")
	}

	return shell(command)
}

// 执行shell命令
func shell(command string) error {
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
