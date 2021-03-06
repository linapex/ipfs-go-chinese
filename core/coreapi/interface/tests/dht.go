
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460167039160320>

package tests

import (
	"context"
	"io"
	"testing"

	"github.com/ipfs/go-ipfs/core/coreapi/interface"
	"github.com/ipfs/go-ipfs/core/coreapi/interface/options"
)

func (tp *provider) TestDht(t *testing.T) {
	tp.hasApi(t, func(api iface.CoreAPI) error {
		if api.Dht() == nil {
			return apiNotImplemented
		}
		return nil
	})

	t.Run("TestDhtFindPeer", tp.TestDhtFindPeer)
	t.Run("TestDhtFindProviders", tp.TestDhtFindProviders)
	t.Run("TestDhtProvide", tp.TestDhtProvide)
}

func (tp *provider) TestDhtFindPeer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apis, err := tp.MakeAPISwarm(ctx, true, 5)
	if err != nil {
		t.Fatal(err)
	}

	self0, err := apis[0].Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	laddrs0, err := apis[0].Swarm().LocalAddrs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(laddrs0) != 1 {
		t.Fatal("unexpected number of local addrs")
	}

	pi, err := apis[2].Dht().FindPeer(ctx, self0.ID())
	if err != nil {
		t.Fatal(err)
	}

	if pi.Addrs[0].String() != laddrs0[0].String() {
		t.Errorf("got unexpected address from FindPeer: %s", pi.Addrs[0].String())
	}

	self2, err := apis[2].Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	pi, err = apis[1].Dht().FindPeer(ctx, self2.ID())
	if err != nil {
		t.Fatal(err)
	}

	laddrs2, err := apis[2].Swarm().LocalAddrs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(laddrs2) != 1 {
		t.Fatal("unexpected number of local addrs")
	}

	if pi.Addrs[0].String() != laddrs2[0].String() {
		t.Errorf("got unexpected address from FindPeer: %s", pi.Addrs[0].String())
	}
}

func (tp *provider) TestDhtFindProviders(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apis, err := tp.MakeAPISwarm(ctx, true, 5)
	if err != nil {
		t.Fatal(err)
	}

	p, err := addTestObject(ctx, apis[0])
	if err != nil {
		t.Fatal(err)
	}

	out, err := apis[2].Dht().FindProviders(ctx, p, options.Dht.NumProviders(1))
	if err != nil {
		t.Fatal(err)
	}

	provider := <-out

	self0, err := apis[0].Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if provider.ID.String() != self0.ID().String() {
		t.Errorf("got wrong provider: %s != %s", provider.ID.String(), self0.ID().String())
	}
}

func (tp *provider) TestDhtProvide(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apis, err := tp.MakeAPISwarm(ctx, true, 5)
	if err != nil {
		t.Fatal(err)
	}

	off0, err := apis[0].WithOptions(options.Api.Offline(true))
	if err != nil {
		t.Fatal(err)
	}

	s, err := off0.Block().Put(ctx, &io.LimitedReader{R: rnd, N: 4092})
	if err != nil {
		t.Fatal(err)
	}

	p := s.Path()

	out, err := apis[2].Dht().FindProviders(ctx, p, options.Dht.NumProviders(1))
	if err != nil {
		t.Fatal(err)
	}

	provider := <-out

	self0, err := apis[0].Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if provider.ID.String() != "<peer.ID >" {
		t.Errorf("got wrong provider: %s != %s", provider.ID.String(), self0.ID().String())
	}

	err = apis[0].Dht().Provide(ctx, p)
	if err != nil {
		t.Fatal(err)
	}

	out, err = apis[2].Dht().FindProviders(ctx, p, options.Dht.NumProviders(1))
	if err != nil {
		t.Fatal(err)
	}

	provider = <-out

	if provider.ID.String() != self0.ID().String() {
		t.Errorf("got wrong provider: %s != %s", provider.ID.String(), self0.ID().String())
	}
}

