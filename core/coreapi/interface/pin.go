
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460166468734976>

package iface

import (
	"context"

	options "github.com/ipfs/go-ipfs/core/coreapi/interface/options"
)

//pin保存有关pinned资源的信息
type Pin interface {
//固定对象的路径
	Path() ResolvedPath

//引脚类型
	Type() string
}

//pinstatus保存有关pin运行状况的信息
type PinStatus interface {
//OK指示是否已验证PIN正确。
	Ok() bool

//bad nodes从pin返回任何坏（通常丢失）节点
	BadNodes() []BadPinNode
}

//bad pin node是一个被pin标记为坏的节点。请验证
type BadPinNode interface {
//路径是节点的路径
	Path() ResolvedPath

//err是将节点标记为坏节点的原因
	Err() error
}

//pinapi指定Pining的接口
type PinAPI interface {
//添加创建新的pin，默认为递归-固定整个引用的pin
//树
	Add(context.Context, Path, ...options.PinAddOption) error

//ls返回此节点上固定对象的列表
	Ls(context.Context, ...options.PinLsOption) ([]Pin, error)

//rm删除由路径指定的对象的pin
	Rm(context.Context, Path) error

//更新将一个管脚更改为另一个管脚，跳过对中匹配路径的检查
//老树
	Update(ctx context.Context, from Path, to Path, opts ...options.PinUpdateOption) error

//验证验证固定对象的完整性
	Verify(context.Context) (<-chan PinStatus, error)
}

