
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:46</date>
//</624460192246927360>

package unit

import "fmt"

type Information int64

const (
_  Information = iota //通过分配给空标识符忽略第一个值
	KB             = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func (i Information) String() string {

	tmp := int64(i)

//违约
	var d = tmp
	symbol := "B"

	switch {
	case i > EB:
		d = tmp / EB
		symbol = "EB"
	case i > PB:
		d = tmp / PB
		symbol = "PB"
	case i > TB:
		d = tmp / TB
		symbol = "TB"
	case i > GB:
		d = tmp / GB
		symbol = "GB"
	case i > MB:
		d = tmp / MB
		symbol = "MB"
	case i > KB:
		d = tmp / KB
		symbol = "KB"
	}
	return fmt.Sprintf("%d %s", d, symbol)
}

