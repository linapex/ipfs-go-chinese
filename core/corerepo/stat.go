
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460172760190976>

package corerepo

import (
	"fmt"
	"math"

	context "context"

	"github.com/ipfs/go-ipfs/core"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"

	humanize "gx/ipfs/QmPSBJL4momYnE7DcUyk2DVhD6rH488ZmHBGLbxNdhU44K/go-humanize"
)

//sizestat包装有关存储库大小及其限制的信息。
type SizeStat struct {
RepoSize   uint64 //字节大小
StorageMax uint64 //字节大小
}

//stat包装有关存储在磁盘上的对象的信息。
type Stat struct {
	SizeStat
	NumObjects uint64
	RepoPath   string
	Version    string
}

//nolimit表示无限存储的值
const NoLimit uint64 = math.MaxUint64

//repostat返回一个设置了所有字段的*stat对象。
func RepoStat(ctx context.Context, n *core.IpfsNode) (Stat, error) {
	sizeStat, err := RepoSize(ctx, n)
	if err != nil {
		return Stat{}, err
	}

	allKeys, err := n.Blockstore.AllKeysChan(ctx)
	if err != nil {
		return Stat{}, err
	}

	count := uint64(0)
	for range allKeys {
		count++
	}

	path, err := fsrepo.BestKnownPath()
	if err != nil {
		return Stat{}, err
	}

	return Stat{
		SizeStat: SizeStat{
			RepoSize:   sizeStat.RepoSize,
			StorageMax: sizeStat.StorageMax,
		},
		NumObjects: count,
		RepoPath:   path,
		Version:    fmt.Sprintf("fs-repo@%d", fsrepo.RepoVersion),
	}, nil
}

//
func RepoSize(ctx context.Context, n *core.IpfsNode) (SizeStat, error) {
	r := n.Repo

	cfg, err := r.Config()
	if err != nil {
		return SizeStat{}, err
	}

	usage, err := r.GetStorageUsage()
	if err != nil {
		return SizeStat{}, err
	}

	storageMax := NoLimit
	if cfg.Datastore.StorageMax != "" {
		storageMax, err = humanize.ParseBytes(cfg.Datastore.StorageMax)
		if err != nil {
			return SizeStat{}, err
		}
	}

	return SizeStat{
		RepoSize:   usage,
		StorageMax: storageMax,
	}, nil
}

