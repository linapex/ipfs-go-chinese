
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:37</date>
//</624460153713856512>

//+生成testrunmain

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//这种滥用非常严重，以至于我在写这段代码时感到很肮脏。
//但这是在不编写自定义编译器的情况下完成此任务的唯一方法
//通过go测试成为go build的克隆
func TestRunMain(t *testing.T) {
	args := flag.Args()
	os.Args = append([]string{os.Args[0]}, args...)
	ret := mainRet()

	p := os.Getenv("IPFS_COVER_RET_FILE")
	if len(p) != 0 {
		ioutil.WriteFile(p, []byte(fmt.Sprintf("%d\n", ret)), 0777)
	}

//关闭输出，这样Go测试就不会打印任何内容
	null, _ := os.Open(os.DevNull)
	os.Stderr = null
	os.Stdout = null
}

