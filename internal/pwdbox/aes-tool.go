package pwdbox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 辅助保存 Key 和 iv 的类
type AesInfo struct {
	Key []byte
	IV  []byte
}

func NewAesInfo(key string, iv string) AesInfo {
	return AesInfo{
		Key: []byte(key),
		IV:  []byte(iv),
	}
}

func TestAES() {

	key := []byte("bGcGfWb3Kg2s4gcG")
	iv := []byte("aebksHkG4jAEk2Ag")

	rawText := "hello world"

	plaintext := []byte(rawText)
	// 加密
	ciphertext, err := Encrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密结果 ：%s\n", base64.StdEncoding.EncodeToString(ciphertext))

	// 解密
	decrypted, err := Decrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密结果 ：%s\n", decrypted)

	ciphertext2, err2 := EncryptToString(rawText, key, iv)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("加密结果2：" + ciphertext2 + "\n")

	// 解密
	decrypted2, err2 := DecryptToString(ciphertext2, key, iv)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("解密结果2：" + decrypted2 + "\n")
}

// 加密 AES-128-CBC
func Encrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = PKCS5Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}
func EncryptToString(plaintext string, key []byte, iv []byte) (string, error) {
	text := ""
	bs, err := Encrypt([]byte(plaintext), key, iv)
	if err == nil {
		text = base64.StdEncoding.EncodeToString(bs)
	}
	return text, err
}

// 解密
func Decrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)

	return plaintext, nil
}
func DecryptToString(ciphertext string, key []byte, iv []byte) (string, error) {
	text := ""
	rawbs, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return text, err
	}

	bs, err := Decrypt(rawbs, key, iv)
	if err == nil {
		text = string(bs)
	}
	return text, err
}

// 填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去除填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
