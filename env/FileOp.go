package env

import (
	"context"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	imgtype "github.com/shamsher31/goimgtype"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 大写开头的变量不符合 js 的风格，用 tag 重新设置字段在 json 中的名字
type CusFileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	// Mode FileMode     // file mode bits
	// ModTime time.Time // modification time
	IsDir bool `json:"isDir"` // abbreviation for Mode().IsDir()
	// Sys   any  // underlying data source (can return nil)

	Path string `json:"path"`
	Ext  string `json:"ext"` // 扩展类型
}

func NewCusFileInfo(info fs.FileInfo, parentPath string) CusFileInfo {
	return CusFileInfo{
		Name:  info.Name(),
		Size:  info.Size(),
		IsDir: info.IsDir(),
		Path:  filepath.Join(parentPath, info.Name()),
		Ext:   filepath.Ext(info.Name()),
	}
}

type FileOp struct {
	ctx context.Context
}

// 获取文件夹下所有文件
// js接收的时候会自动转换成 json 对象
func (fo *FileOp) List(path string) []CusFileInfo {
	var list = []CusFileInfo{}
	if strings.Trim(path, " ") == "" {
		return list
	}

	fi, err := os.Stat(path)
	// if err != nil && os.IsNotExist(err) {
	if err != nil {
		return list
	}

	if !fi.IsDir() {
		return list
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("路径读取失败！", err)
		log.Panicln(err)
		return list
	}

	for _, f := range files {
		list = append(list, NewCusFileInfo(f, path))
		// list[index] = NewCusFileInfo(f, path)
	}

	return list
}

// 获取文件夹下所有图片文件
func (fo *FileOp) ListImage(path string) []CusFileInfo {
	var list = fo.List(path)
	var re = []CusFileInfo{}

	for _, cfi := range list {
		if cfi.IsDir || cfi.Size <= 0 {
			continue
		}

		var datatype string
		var err any

		func() {
			defer func() {
				err = recover()
			}()
			// 为什么没有处理空文件啊！
			datatype, err = imgtype.Get(cfi.Path)
		}()

		if err == nil && strings.HasPrefix(datatype, `image/`) {
			re = append(re, cfi)
		}
	}

	return re
}

// 读取文件 （返回的字节数组会自动base64编码成字符串）
func (fo *FileOp) Open(path string) []byte {
	_, err := os.Stat(path)
	// if err != nil && os.IsNotExist(err) {
	if err != nil {
		log.Println(err)
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println("文件打开失败")
		log.Println(err)
	}
	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Println("文件读取失败")
		log.Println(err2)
	}

	return data
}

func (fo *FileOp) SetContext(ctx context.Context) {
	fo.ctx = ctx
}

// 打开文件夹选择
func (fo *FileOp) SelectFolder(title string) string {
	// var p = GetCurrentAbPath()
	folder, err := runtime.OpenDirectoryDialog(fo.ctx, runtime.OpenDialogOptions{
		// DefaultDirectory: p,
	})
	// return nil
	if err != nil {
		log.Println(err)
		return ""
	}
	return folder
}
