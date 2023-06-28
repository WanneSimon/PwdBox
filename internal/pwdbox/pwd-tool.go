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

	// #key 一起存储的方案 （key 和 verifyData 一起存）
	// fileContent := key + "\n\r" + verifyData
	fileContent := verifyData
	data, err := EncryptToString(fileContent, []byte(key), []byte(iv))
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
		return false, errors.New("未找到验证文件")
	}

	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Println("验证文件读取失败2")
		log.Println(err2)
		return false, errors.New("验证文件读取失败")
	}

	fileContent := string(data)
	// arr := strings.Split(fileContent, "\n\r")
	// #key 一起存储的方案
	// if !strings.HasPrefix(fileContent, key) {
	// 	log.Println("验证文件读取失败2")
	// 	log.Println(err2)
	// 	return false, errors.New("key错误")
	// }
	// log.Println("===data===\n", str)
	// 解密后比较原数据是否相等（key错误时解密会报错）
	// rawStr, err3 := DecryptToString(fileContent, []byte(key), []byte(iv))
	// log.Println("===data===\n", rawStr, verifyData)
	// if err3 == nil && rawStr == inputContent { // 比较解密字符串

	// 加密后比较是否相等
	// inputContent := key + "\n\r" + verifyData
	inputContent := verifyData
	encStr, err3 := EncryptToString(inputContent, []byte(key), []byte(iv))
	if err3 == nil && encStr == fileContent { // 比较加密字符串
		aesHolder.Key = []byte(key)
		aesHolder.IV = []byte(iv)
		return true, nil
	} else if err3 != nil {
		log.Println("验证数据失败")
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
