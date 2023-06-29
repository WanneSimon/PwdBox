package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	dataop "github.com/wanneSimon/pwdbox/internal/data-op"
	"github.com/wanneSimon/pwdbox/internal/pwdbox"
)

func main0() {
	str := time.Now().Format("2006-01-02")
	fmt.Println(str)

	uri := "config/pwdbox.db3"
	key := "1234567890abcdef"
	iv := "1234567890abcdef"

	pwdbox.InitSqlite(uri)

	op := dataop.DataOutOp{}
	txt, err := op.ExportToMarkdown([]byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("text.md", []byte(txt), 0666)
	if err != nil {
		panic(err)
	}

	log.Println("finished")
}
