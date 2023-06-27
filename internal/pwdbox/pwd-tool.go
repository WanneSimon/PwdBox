package pwdbox

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// 持有并验证用户输入的 key 和 iv 是否正确
var aesHolder AesInfo

const verifyData string = "this-is-check-data"
const verifyDataPath = "config/key-iv-for-check"

// 加密解密工具类
type PwdTool struct {
}

func checkParameter(key string, iv string) error {
	if len(key) != 16 {
		return errors.New("key不是16位")
	}
	if len(iv) != 16 {
		return errors.New("iv不是16位")
	}
	return nil
}

// 检查是否存在验证数据
func (pt *PwdTool) ExistVerifyData() bool {
	_, err := os.Open(verifyDataPath)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 保存验证数据
func (pt *PwdTool) SaveAesInfo(key string, iv string) bool {
	checkRe := checkParameter(key, iv)
	if checkRe != nil {
		log.Println(checkRe)
		return false
	}

	data, err := EncryptToString(verifyData, []byte(key), []byte(iv))
	if err != nil {
		log.Println("验证数据加密失败")
		log.Println(err)
		return false
	}

	parentPath := filepath.Dir(verifyDataPath)
	if parentPath != "." {
		os.MkdirAll(parentPath, 0766)
	}

	err2 := ioutil.WriteFile(verifyDataPath, []byte(data), 0666)
	if err2 != nil {
		log.Println("验证文件保存失败")
		log.Println(err2)
		return false
	}

	aesHolder.Key = []byte(key)
	aesHolder.IV = []byte(iv)
	return true
}

// 读取验证文件，然后验证 key 和 iv 是否正确
func (pt *PwdTool) VerifyAndKeepAesInfo(key string, iv string) (bool, error) {
	checkRe := checkParameter(key, iv)
	if checkRe != nil {
		log.Println(checkRe)
		return false, checkRe
	}

	file, err := os.Open(verifyDataPath)
	if err != nil {
		log.Println("未找到验证文件2")
		log.Println(err)
		return false, err
	}

	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Println("验证文件读取失败2")
		log.Println(err2)
	}

	str := string(data)
	log.Println("===data===\n", str)
	rawStr, err3 := DecryptToString(str, []byte(key), []byte(iv))
	log.Println("===data===\n", rawStr, verifyData)
	if err3 == nil && rawStr == verifyData {
		aesHolder.Key = []byte(key)
		aesHolder.IV = []byte(iv)
		return true, nil
	} else if err3 != nil {
		log.Println("验证数据解密失败")
		log.Println(err3)
	}
	return false, errors.New("验证失败")
}

// 该方法只是为了让 js 中有 AesInfo 这个类型
func (pt *PwdTool) GetAesInfo() AesInfo {
	return aesHolder
}

// 解密方法
func (pt *PwdTool) DecryptPwd(encPwd string) string {
	re, err := DecryptToString(encPwd, aesHolder.Key, aesHolder.IV)
	if err != nil {
		return ""
	}
	return re
}
