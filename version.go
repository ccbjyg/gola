package main

import (
	"flag"
	"fmt"
)

//flags
var (
	debug = flag.Bool("debug", false, "打开以输出调试信息")
)

var (
	version = "version: 0.0.1"
)

//Parse 解析flag 及版本信息
func Parse() {
	flag.Parse()
	fmt.Println(version)
	fmt.Println("支持如下基础操作  + - * / ")
	fmt.Println("内置如下方法 max(n1,n2,n3...) min(n1,n2,n3...) lg(n) ln(n)")
	fmt.Println("--------------------------------")
}
