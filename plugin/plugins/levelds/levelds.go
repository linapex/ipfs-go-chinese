
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460183195619328>

package levelds

import (
	"fmt"
	"path/filepath"

	"github.com/ipfs/go-ipfs/plugin"
	"github.com/ipfs/go-ipfs/repo"
	"github.com/ipfs/go-ipfs/repo/fsrepo"

	levelds "gx/ipfs/QmUhiHo586S2XpAFvkL1xDxeNwVHVQg7sDTxzS8ituQawr/go-ds-leveldb"
	ldbopts "gx/ipfs/QmbBhyDKsY4mbY6xsKt3qu9Y7FPvMJ6qbD8AMjYYvPRw1g/goleveldb/leveldb/opt"
)

//插件导出将加载的插件列表
var Plugins = []plugin.Plugin{
	&leveldsPlugin{},
}

type leveldsPlugin struct{}

var _ plugin.PluginDatastore = (*leveldsPlugin)(nil)

func (*leveldsPlugin) Name() string {
	return "ds-level"
}

func (*leveldsPlugin) Version() string {
	return "0.1.0"
}

func (*leveldsPlugin) Init() error {
	return nil
}

func (*leveldsPlugin) DatastoreTypeName() string {
	return "levelds"
}

type datastoreConfig struct {
	path        string
	compression ldbopts.Compression
}

//badgerdsdatastoreconfig返回badger数据存储区的配置存根
//从给定的参数
func (*leveldsPlugin) DatastoreConfigParser() fsrepo.ConfigFromMap {
	return func(params map[string]interface{}) (fsrepo.DatastoreConfig, error) {
		var c datastoreConfig
		var ok bool

		c.path, ok = params["path"].(string)
		if !ok {
			return nil, fmt.Errorf("'path' field is missing or not string")
		}

		switch cm := params["compression"].(string); cm {
		case "none":
			c.compression = ldbopts.NoCompression
		case "snappy":
			c.compression = ldbopts.SnappyCompression
		case "":
			c.compression = ldbopts.DefaultCompression
		default:
			return nil, fmt.Errorf("unrecognized value for compression: %s", cm)
		}

		return &c, nil
	}
}

func (c *datastoreConfig) DiskSpec() fsrepo.DiskSpec {
	return map[string]interface{}{
		"type": "levelds",
		"path": c.path,
	}
}

func (c *datastoreConfig) Create(path string) (repo.Datastore, error) {
	p := c.path
	if !filepath.IsAbs(p) {
		p = filepath.Join(path, p)
	}

	return levelds.NewDatastore(p, &levelds.Options{
		Compression: c.compression,
	})
}

