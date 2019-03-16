
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:42</date>
//</624460176635727872>

//+构建Linux darwin freebsd netbsd openbsd
//+建设！诺福斯

package ipns

import (
	core "github.com/ipfs/go-ipfs/core"
	mount "github.com/ipfs/go-ipfs/fuse/mount"
)

//mount在给定位置安装IPN，并返回mount.mount实例。
func Mount(ipfs *core.IpfsNode, ipnsmp, ipfsmp string) (mount.Mount, error) {
	cfg, err := ipfs.Repo.Config()
	if err != nil {
		return nil, err
	}

	allow_other := cfg.Mounts.FuseAllowOther

	fsys, err := NewFileSystem(ipfs, ipfs.PrivateKey, ipfsmp, ipnsmp)
	if err != nil {
		return nil, err
	}

	return mount.NewMount(ipfs.Process(), fsys, ipnsmp, allow_other)
}

