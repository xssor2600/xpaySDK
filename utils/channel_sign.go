package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	mrand "math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func KsSignFromMap(s map[string]interface{}, key string) string {
	sorted_keys := make([]string, 0)
	for k := range s {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", s[k])
		if k != "" && value != "" && k != "sign" && k != "-" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if key != "" && len(signStrings) > 0 {
		signStrings = signStrings[0:len(signStrings)-1] + key
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := hex.EncodeToString(cipherStr)
	return upperSign
}

func GenChannelSign(method, url, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (string, error) {
	//method内容必须大写，如GET、POST，uri不包含域名，必须以'/'开头
	targetStr := method + "\n" + url + "\n" + timestamp + "\n" + nonce + "\n" + body + "\n"
	h := sha256.New()
	h.Write([]byte(targetStr))
	digestBytes := h.Sum(nil)

	signBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digestBytes)
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)

	return sign, nil
}

func CheckChannelSign(timestamp, nonce, body, signature, pubKeyStr string) error {

	pubKey, err := PemToRSAPublicKey(pubKeyStr)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256([]byte(timestamp + "\n" + nonce + "\n" + body + "\n"))
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	return err
}

func PemToRSAPublicKey(pemKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemKeyStr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, fmt.Errorf("empty block in pem string")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, fmt.Errorf("not rsa public key")
	}
}

func GetByteAuth(appId, onceStr string, timestamp string, keyVersion string, signature string) string {
	return fmt.Sprintf("SHA256-RSA2048 appid=%s,nonce_str=%s,key_version=%s,timestamp=%s,signature=%s",
		appId, onceStr, keyVersion, timestamp, signature)
}

// GetNonceStr 32位随机字符串
func GetNonceStr() string {
	return strings.ToUpper(RandString(32))
}

func GetTimeStamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func init() {
	mrand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("1112222222")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mrand.Intn(len(letterRunes))]
	}
	return string(b)
}
