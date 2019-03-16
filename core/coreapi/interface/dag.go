
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460165155917824>

package iface

import (
	"context"
	"io"

	"github.com/ipfs/go-ipfs/core/coreapi/interface/options"

	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

//DAGOPS对可以一起批处理的操作进行分组
type DagOps interface {
//使用指定的格式和输入编码放置插入数据。
//除非与withcodec或withhash一起使用，否则默认值为“dag cbor”和
//使用“sha256”。
	Put(ctx context.Context, src io.Reader, opts ...options.DagPutOption) (ResolvedPath, error)
}

//dagbatch是dagapi的批处理版本。dagbatch的所有实现
//应该是螺纹安全的
type DagBatch interface {
	DagOps

//提交将节点提交到数据存储并向网络公布它们
	Commit(ctx context.Context) error
}

//dagapi指定到ipld的接口
type DagAPI interface {
	DagOps

//获取解析和获取由路径指定的节点的尝试
	Get(ctx context.Context, path Path) (ipld.Node, error)

//树返回由路径指定的节点内的路径列表。
	Tree(ctx context.Context, path Path, opts ...options.DagTreeOption) ([]Path, error)

//批量创建新的dagbatch
	Batch(ctx context.Context) DagBatch
}

