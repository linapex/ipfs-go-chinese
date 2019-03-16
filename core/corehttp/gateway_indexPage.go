
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460171556425728>

package corehttp

import (
	"html/template"
	"net/url"
	"path"
	"strings"

	"github.com/ipfs/go-ipfs/assets"
)

//目录列表结构
type listingTemplateData struct {
	Listing  []directoryItem
	Path     string
	BackLink string
	Hash     string
}

type directoryItem struct {
	Size string
	Name string
	Path string
}

var listingTemplate *template.Template

func init() {
	knownIconsBytes, err := assets.Asset("dir-index-html/knownIcons.txt")
	if err != nil {
		panic(err)
	}
	knownIcons := make(map[string]struct{})
	for _, ext := range strings.Split(strings.TrimSuffix(string(knownIconsBytes), "\n"), "\n") {
		knownIcons[ext] = struct{}{}
	}

//通过扩展名猜测其类型/图标的帮助程序
	iconFromExt := func(name string) string {
		ext := path.Ext(name)
		_, ok := knownIcons[ext]
		if !ok {
//默认空白图标
			return "ipfs-_blank"
		}
return "ipfs-" + ext[1:] //第一个点的切片
	}

//自定义模板转义函数以转义完整路径，包括“”和“？”
	urlEscape := func(rawUrl string) string {
		pathUrl := url.URL{Path: rawUrl}
		return pathUrl.String()
	}

//目录列表模板
	dirIndexBytes, err := assets.Asset("dir-index-html/dir-index.html")
	if err != nil {
		panic(err)
	}

	listingTemplate = template.Must(template.New("dir").Funcs(template.FuncMap{
		"iconFromExt": iconFromExt,
		"urlEscape":   urlEscape,
	}).Parse(string(dirIndexBytes)))
}

