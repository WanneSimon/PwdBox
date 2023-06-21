package main

import (
	"fmt"
	"time"
)

func main0() {
	str := time.Now().Format("2006-01-02")
	fmt.Println(str)
}
