package dataop

import (
	"bytes"
	_ "embed"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/wanneSimon/pwdbox/internal/pwdbox"
)

type MdVo struct {
	Plat     pwdbox.Platform
	Accounts []pwdbox.Account
}

// var DefaultMdTpl embed.FS

//go:embed md.tpl
var DefaultMdTpl string

func (a *MdVo) DefaultToMarkdownText() (string, error) {
	// log.Println("DefaultMdTpl", DefaultMdTpl)
	// tmp, err := template.ParseFS(DefaultMdTpl)
	tmp, err := template.New("").Parse(DefaultMdTpl)
	if err != nil {
		log.Println(err)
		return "", errors.New("模版读取失败！")
	}
	return a.toText(tmp), nil
}

func (a *MdVo) ToMarkdownText(filename string) (string, error) {
	tmp, err := template.ParseFiles(filename)
	if err != nil {
		return "", err
	}
	// template.Must()
	// buf := new(bytes.Buffer)
	// tmp.ExecuteTemplate(buf, "PlatformTemplate", a)
	// return buf.String(), nil
	return a.toText(tmp), nil
}

func (a *MdVo) toText(tmp *template.Template) string {
	buf := new(bytes.Buffer)
	tmp.ExecuteTemplate(buf, "PlatformTemplate", a)
	return buf.String()
}

type DataOutOp struct {
}

// 检查导出的文件是否已存在, 存在则返回文件的绝对路径
func (do *DataOutOp) ExportFileExist() string {
	file := "pwdbox-export.md"
	_, err := os.Stat(file)

	if err == nil {
		absPath, _ := filepath.Abs(file)
		return absPath
	}

	if os.IsNotExist(err) {
		return ""
	}

	return ""
}

// 对外导出方法，成功则返回文件路径
func (do *DataOutOp) ExportAllToMarkdown() (string, error) {
	txt, err := do.ExportToMarkdown(pwdbox.AesHolder.Key, pwdbox.AesHolder.IV)
	if err != nil {
		log.Println(err)
		return "", err
	}

	file := "pwdbox-export.md"
	absPath, _ := filepath.Abs(file)
	err = ioutil.WriteFile(absPath, []byte(txt), 0666)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return absPath, nil
}

// 导出所有账号
func (do *DataOutOp) ExportToMarkdown(key []byte, iv []byte) (string, error) {
	page := 0
	pageSize := 500
	total := 0

	finalText := ""

	for page == 0 || (page*pageSize) < total {
		page = page + 1
		// 遍历Platform
		var data pwdbox.PageData[pwdbox.Platform] = pwdbox.PlatformServiceInstance.PageList("", page, pageSize)
		total = data.Total

		if total == 0 {
			break
		}

		pList := data.Data
		for _, plat := range pList {
			// 获取Platform 的 Account
			accountList := pwdbox.AccountServiceInstance.List(plat.Id, "", "", "", 1, 999999)

			// 密码解码
			for i := 0; i < len(accountList); i++ {
				item := &accountList[i]
				re, err := pwdbox.DecryptToString(item.Password, key, iv)
				if err != nil {
					return "", err
				}
				item.Password = re
			}

			mdvo := MdVo{
				Plat:     plat,
				Accounts: accountList,
			}
			// 转 md 文本
			mdtxt, err := mdvo.DefaultToMarkdownText()
			if err != nil {
				return "", err
			}
			finalText = finalText + mdtxt
		}
	}

	return finalText, nil
}
