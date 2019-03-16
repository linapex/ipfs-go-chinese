
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:45</date>
//</624460186936938496>

package integrationtest

import (
	"bytes"
	"context"
	"testing"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/mock"
	"gx/ipfs/QmWoXtvgC8inqFkAATB7cp2Dax7XBi9VDvSg9RCCZufmRk/go-block-format"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	mocknet "gx/ipfs/QmYxivS34F2M2n44WQQnRHGAKS8aoRUxwGpi9wk4Cdn4Jf/go-libp2p/p2p/net/mock"
)

func TestBitswapWithoutRouting(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const numPeers = 4

//创建网络
	mn := mocknet.New(ctx)

	var nodes []*core.IpfsNode
	for i := 0; i < numPeers; i++ {
		n, err := core.NewNode(ctx, &core.BuildCfg{
			Online:  true,
			Host:    coremock.MockHostOption(mn),
Routing: core.NilRouterOption, //无路由
		})
		if err != nil {
			t.Fatal(err)
		}
		defer n.Close()
		nodes = append(nodes, n)
	}

	mn.LinkAll()

//连接它们
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1 == n2 {
				continue
			}

			log.Debug("connecting to other hosts")
			p2 := n2.PeerHost.Peerstore().PeerInfo(n2.PeerHost.ID())
			if err := n1.PeerHost.Connect(ctx, p2); err != nil {
				t.Fatal(err)
			}
		}
	}

//向前面的每个添加块
	log.Debug("adding block.")
	block0 := blocks.NewBlock([]byte("block0"))
	block1 := blocks.NewBlock([]byte("block1"))

//先放1
	if err := nodes[0].Blockstore.Put(block0); err != nil {
		t.Fatal(err)
	}

//把它拿出来。
	for i, n := range nodes {
//先跳过，因为块不在其交换中。将绞死。
		if i == 0 {
			continue
		}

		log.Debugf("%d %s get block.", i, n.Identity)
		b, err := n.Blocks.GetBlock(ctx, cid.NewCidV0(block0.Multihash()))
		if err != nil {
			t.Error(err)
		} else if !bytes.Equal(b.RawData(), block0.RawData()) {
			t.Error("byte comparison fail")
		} else {
			log.Debug("got block: %s", b.Cid())
		}
	}

//后放1
	if err := nodes[1].Blockstore.Put(block1); err != nil {
		t.Fatal(err)
	}

//把它拿出来。
	for _, n := range nodes {
		b, err := n.Blocks.GetBlock(ctx, cid.NewCidV0(block1.Multihash()))
		if err != nil {
			t.Error(err)
		} else if !bytes.Equal(b.RawData(), block1.RawData()) {
			t.Error("byte comparison fail")
		} else {
			log.Debug("got block: %s", b.Cid())
		}
	}
}

