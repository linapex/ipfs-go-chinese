
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460182847492096>

package plugin

//插件是各种go ipfs插件的基本接口
//它将包含在不同插件的接口中
type Plugin interface {
//名称应返回插件的唯一名称
	Name() string
//版本返回插件的当前版本
	Version() string
//加载插件时调用一次init
	Init() error
}

