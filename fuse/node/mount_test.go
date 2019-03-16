
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:42</date>
//</624460177139044352>

//+建设！诺福斯

package node

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"context"

	core "github.com/ipfs/go-ipfs/core"
	ipns "github.com/ipfs/go-ipfs/fuse/ipns"
	mount "github.com/ipfs/go-ipfs/fuse/mount"

	ci "gx/ipfs/QmNvHv84aH2qZafDuSdKJCQ1cvPZ1kmQmyD4YtzjUHuk9v/go-testutil/ci"
)

func maybeSkipFuseTests(t *testing.T) {
	if ci.NoFuse() {
		t.Skip("Skipping FUSE tests")
	}
}

func mkdir(t *testing.T, path string) {
	err := os.Mkdir(path, os.ModeDir|os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

//测试外部卸载，然后尝试在代码中卸载
func TestExternalUnmount(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

//待办事项：需要吗？
	maybeSkipFuseTests(t)

	node, err := core.NewNode(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	err = ipns.InitializeKeyspace(node, node.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

//获取测试目录路径（/tmp/fusetestxxxx）
	dir, err := ioutil.TempDir("", "fusetest")
	if err != nil {
		t.Fatal(err)
	}

	ipfsDir := dir + "/ipfs"
	ipnsDir := dir + "/ipns"
	mkdir(t, ipfsDir)
	mkdir(t, ipnsDir)

	err = Mount(node, ipfsDir, ipnsDir)
	if err != nil {
		t.Fatal(err)
	}

//运行shell命令从外部卸载目录
	cmd, err := mount.UnmountCmd(ipfsDir)
	if err != nil {
		t.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

//托多（noffle）：运行fs的Goroutine需要一段时间。要得到通知并进行清理。
	time.Sleep(time.Millisecond * 100)

//尝试卸载IPF；应成功卸载。
	err = node.Mounts.Ipfs.Unmount()
	if err != mount.ErrNotMounted {
		t.Fatal("Unmount should have failed")
	}
}

