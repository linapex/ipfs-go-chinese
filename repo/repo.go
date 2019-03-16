
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460184663625728>

package repo

import (
	"errors"
	"io"

	filestore "github.com/ipfs/go-ipfs/filestore"
	keystore "github.com/ipfs/go-ipfs/keystore"

	ma "gx/ipfs/QmNTCey11oxhb1AxDnQBRHtdhap6Ctud872NjAYPYYXPuc/go-multiaddr"
	config "gx/ipfs/QmcRKBUqc2p3L1ZraoJjbXfs9E6xzvEuyK9iypb5RGwfsr/go-ipfs-config"
	ds "gx/ipfs/Qmf4xQhNomPNhrtZc67qSnfJSjxjXs9LWvknJtSXwimPrM/go-datastore"
)

var (
	ErrApiNotRunning = errors.New("api not running")
)

//repo表示给定ipfs节点的所有持久数据。
type Repo interface {
//config从repo返回ipfs配置文件。作出改变
//返回的配置不会自动保留。
	Config() (*config.Config, error)

//backupconfig使用创建当前配置文件的备份
//用于命名的给定前缀。
	BackupConfig(prefix string) (string, error)

//setconfig将给定的配置结构保存到存储中。
	SetConfig(*config.Config) error

//setconfigkey在config中设置给定的键值对，并将其保存到存储器中。
	SetConfigKey(key string, value interface{}) error

//getconfigkey从存储中的配置中读取给定密钥的值。
	GetConfigKey(key string) (interface{}, error)

//数据存储返回对配置的数据存储后端的引用。
	Datastore() Datastore

//GetStorageUsage返回存储的字节数。
	GetStorageUsage() (uint64, error)

//keystore返回对密钥管理接口的引用。
	Keystore() keystore.Keystore

//文件管理器返回对文件存储文件管理器的引用。
	FileManager() *filestore.FileManager

//setapiaddr设置repo中的api地址。
	SetAPIAddr(addr ma.Multiaddr) error

//swarmkey返回为专用网络功能配置的共享对称密钥。
	SwarmKey() ([]byte, error)

	io.Closer
}

//数据存储是数据存储需要的接口
//FSrepo可接受。
type Datastore interface {
ds.Batching //应该是安全的，小心点
	io.Closer
}

