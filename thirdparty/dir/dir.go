
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:46</date>
//</624460191890411520>

package dir

//TODO移动到一般位置

import (
	"errors"
	"os"
	"path/filepath"
)

//可写确保目录存在且可写
func Writable(path string) error {
//如果缺少，则构造路径
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
//检查目录是否可写
	if f, err := os.Create(filepath.Join(path, "._check_writable")); err == nil {
		f.Close()
		os.Remove(f.Name())
	} else {
		return errors.New("'" + path + "' is not writable")
	}
	return nil
}

