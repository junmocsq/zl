package main

import (
	"os"
	"wangqingshui/library"
)

func main() {
	library.RegisterConfig(os.Getenv("HOME") + "/www/wangqingshui")
	library.NewConfig()
}
