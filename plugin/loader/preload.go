
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460182713274368>

package loader

import (
	"github.com/ipfs/go-ipfs/plugin"
	pluginbadgerds "github.com/ipfs/go-ipfs/plugin/plugins/badgerds"
	pluginflatfs "github.com/ipfs/go-ipfs/plugin/plugins/flatfs"
	pluginipldgit "github.com/ipfs/go-ipfs/plugin/plugins/git"
	pluginlevelds "github.com/ipfs/go-ipfs/plugin/plugins/levelds"
)

//不编辑此文件
//此文件正在作为插件生成过程的一部分生成
//要更改它，请修改plugin/loader/preload.sh

var preloadPlugins = []plugin.Plugin{
	pluginipldgit.Plugins[0],
	pluginbadgerds.Plugins[0],
	pluginflatfs.Plugins[0],
	pluginlevelds.Plugins[0],
}

