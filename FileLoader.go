package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// var err error
	path := strings.TrimPrefix(req.URL.Path, "/")

	// 处理访问文件（危险操作！有安全问题）
	if strings.HasPrefix(path, "file/") {
		filePath := strings.TrimPrefix(path, "file/")

		println("[Load file]:", filePath)
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(fmt.Sprintf("Could not load file %s", filePath)))
		}

		res.Write(fileData)
	}
}
