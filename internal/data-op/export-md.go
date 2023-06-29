package dataop

import (
	"bytes"
	_ "embed"
	"errors"
	"io/ioutil"
	"log"
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

// 对外导出方法
func (do *DataOutOp) ExportAllToMarkdown() error {
	txt, err := do.ExportToMarkdown(pwdbox.AesHolder.Key, pwdbox.AesHolder.IV)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile("pwdbox-data.md", []byte(txt), 0666)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
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
			for _, item := range accountList {
				// item.Password = tool.DecryptPwd(item.Password)
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
