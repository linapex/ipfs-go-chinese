
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460182398701568>

package plugin

import (
	"github.com/ipfs/go-ipfs/core/coredag"

	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

//pluginipld是一个接口，可以实现它来添加
//对于不同的IPLD格式
type PluginIPLD interface {
	Plugin

	RegisterBlockDecoders(dec ipld.BlockDecoder) error
	RegisterInputEncParsers(iec coredag.InputEncParsers) error
}

