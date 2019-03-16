
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460169094369280>

package iface

import (
	"context"

	"github.com/ipfs/go-ipfs/core/coreapi/interface/options"

	files "gx/ipfs/QmXWZCd8jfaHmt4UDSnjKmGcrQMw95bDGWqEeVLVJjoANX/go-ipfs-files"
	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

type AddEvent struct {
	Name  string
	Path  ResolvedPath `json:",omitempty"`
	Bytes int64        `json:",omitempty"`
	Size  string       `json:",omitempty"`
}

//unixfsapi是IPF中不可变文件的基本接口
//注意：这个API是大量WIP，保证经常中断。
type UnixfsAPI interface {
//添加将读卡器中的数据导入到merkledag文件中
//
//TODO：关于如何在许多不同的场景中使用它的一个长期有用的注释
	Add(context.Context, files.Node, ...options.UnixfsAddOption) (ResolvedPath, error)

//get返回路径引用的文件树的只读句柄
//
//请注意，此API的某些实现可能应用指定的上下文
//到对返回文件执行的操作
	Get(context.Context, Path) (files.Node, error)

//ls返回目录中的链接列表
	Ls(context.Context, Path) ([]*ipld.Link, error)
}

