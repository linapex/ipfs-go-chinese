
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460165025894400>

package iface

import (
	"context"
	"io"

	options "github.com/ipfs/go-ipfs/core/coreapi/interface/options"
)

//blockstat包含有关块的信息
type BlockStat interface {
//大小是块的大小
	Size() int

//path返回块的路径
	Path() ResolvedPath
}

//blockapi指定块层的接口
type BlockAPI interface {
//Put导入原始块数据，使用指定的设置对其进行哈希处理。
	Put(context.Context, io.Reader, ...options.BlockPutOption) (BlockStat, error)

//获取解析路径的尝试并返回块中数据的读卡器
	Get(context.Context, Path) (io.Reader, error)

//RM从本地BlockStore中删除由路径指定的块。
//默认情况下，如果在本地找不到块，将返回错误。
//
//注意：如果固定了指定的块，它将不会被删除，也不会出错。
//将被退回
	Rm(context.Context, Path, ...options.BlockRmOption) error

//stat返回信息
	Stat(context.Context, Path) (BlockStat, error)
}

