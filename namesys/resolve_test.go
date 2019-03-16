
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:43</date>
//</624460179643043840>

package namesys

import (
	"context"
	"errors"
	"testing"
	"time"

	path "gx/ipfs/QmNYPETsdAu2uQ1k9q9S1jYEGURaLHV6cbYRSVFVRftpF8/go-path"

	testutil "gx/ipfs/QmNvHv84aH2qZafDuSdKJCQ1cvPZ1kmQmyD4YtzjUHuk9v/go-testutil"
	mockrouting "gx/ipfs/QmVZ6cQXHoTQja4oo9GhhHZi7dThi4x98mRKgGtKnTy37u/go-ipfs-routing/mock"
	ipns "gx/ipfs/QmWPFehHmySCdaGttQ48iwF7M6mBRrGE5GSPWKCuMWqJDR/go-ipns"
	peer "gx/ipfs/QmY5Grm8pJdiSSVsYxx4uNRgweY72EmYwuSDbRnbFok3iY/go-libp2p-peer"
	ds "gx/ipfs/Qmf4xQhNomPNhrtZc67qSnfJSjxjXs9LWvknJtSXwimPrM/go-datastore"
	dssync "gx/ipfs/Qmf4xQhNomPNhrtZc67qSnfJSjxjXs9LWvknJtSXwimPrM/go-datastore/sync"
)

func TestRoutingResolve(t *testing.T) {
	dstore := dssync.MutexWrap(ds.NewMapDatastore())
	serv := mockrouting.NewServer()
	id := testutil.RandIdentityOrFatal(t)
	d := serv.ClientWithDatastore(context.Background(), id, dstore)

	resolver := NewIpnsResolver(d)
	publisher := NewIpnsPublisher(d, dstore)

	privk, pubk, err := testutil.RandTestKeyPair(512)
	if err != nil {
		t.Fatal(err)
	}

	h := path.FromString("/ipfs/QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN")
	err = publisher.Publish(context.Background(), privk, h)
	if err != nil {
		t.Fatal(err)
	}

	pid, err := peer.IDFromPublicKey(pubk)
	if err != nil {
		t.Fatal(err)
	}

	res, err := resolver.Resolve(context.Background(), pid.Pretty())
	if err != nil {
		t.Fatal(err)
	}

	if res != h {
		t.Fatal("Got back incorrect value.")
	}
}

func TestPrexistingExpiredRecord(t *testing.T) {
	dstore := dssync.MutexWrap(ds.NewMapDatastore())
	d := mockrouting.NewServer().ClientWithDatastore(context.Background(), testutil.RandIdentityOrFatal(t), dstore)

	resolver := NewIpnsResolver(d)
	publisher := NewIpnsPublisher(d, dstore)

	privk, pubk, err := testutil.RandTestKeyPair(512)
	if err != nil {
		t.Fatal(err)
	}

	id, err := peer.IDFromPublicKey(pubk)
	if err != nil {
		t.Fatal(err)
	}

//创建一个过期的记录并将其放入数据存储中
	h := path.FromString("/ipfs/QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN")
	eol := time.Now().Add(time.Hour * -1)

	entry, err := ipns.Create(privk, []byte(h), 0, eol)
	if err != nil {
		t.Fatal(err)
	}
	err = PutRecordToRouting(context.Background(), d, pubk, entry)
	if err != nil {
		t.Fatal(err)
	}

//现在，在系统中已经有一个旧记录的情况下，尝试发布一个新记录
	err = publisher.Publish(context.Background(), privk, h)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyCanResolve(resolver, id.Pretty(), h)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrexistingRecord(t *testing.T) {
	dstore := dssync.MutexWrap(ds.NewMapDatastore())
	d := mockrouting.NewServer().ClientWithDatastore(context.Background(), testutil.RandIdentityOrFatal(t), dstore)

	resolver := NewIpnsResolver(d)
	publisher := NewIpnsPublisher(d, dstore)

	privk, pubk, err := testutil.RandTestKeyPair(512)
	if err != nil {
		t.Fatal(err)
	}

	id, err := peer.IDFromPublicKey(pubk)
	if err != nil {
		t.Fatal(err)
	}

//做好记录，放在数据仓库里
	h := path.FromString("/ipfs/QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN")
	eol := time.Now().Add(time.Hour)
	entry, err := ipns.Create(privk, []byte(h), 0, eol)
	if err != nil {
		t.Fatal(err)
	}
	err = PutRecordToRouting(context.Background(), d, pubk, entry)
	if err != nil {
		t.Fatal(err)
	}

//现在，在系统中已经有一个旧记录的情况下，尝试发布一个新记录
	err = publisher.Publish(context.Background(), privk, h)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyCanResolve(resolver, id.Pretty(), h)
	if err != nil {
		t.Fatal(err)
	}
}

func verifyCanResolve(r Resolver, name string, exp path.Path) error {
	res, err := r.Resolve(context.Background(), name)
	if err != nil {
		return err
	}

	if res != exp {
		return errors.New("got back wrong record")
	}

	return nil
}

