
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460170851782656>

package coredag

import (
	"io"
	"io/ioutil"
	"math"

	"gx/ipfs/QmTQdH4848iTVCJmKXYyRiK72HufWTLYQQ8iN3JaQ8K1Hq/go-merkledag"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
	mh "gx/ipfs/QmerPMzPk1mJVowm8KgmoknWa4yCYvvugMPsgWmDNUvDLW/go-multihash"
)

func dagpbJSONParser(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nd := &merkledag.ProtoNode{}

	err = nd.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	nd.SetCidBuilder(cidPrefix(mhType, mhLen))

	return []ipld.Node{nd}, nil
}

func dagpbRawParser(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nd, err := merkledag.DecodeProtobuf(data)
	if err != nil {
		return nil, err
	}

	nd.SetCidBuilder(cidPrefix(mhType, mhLen))

	return []ipld.Node{nd}, nil
}

func cidPrefix(mhType uint64, mhLen int) *cid.Prefix {
	if mhType == math.MaxUint64 {
		mhType = mh.SHA2_256
	}

	prefix := &cid.Prefix{
		MhType:   mhType,
		MhLength: mhLen,
		Version:  1,
		Codec:    cid.DagProtobuf,
	}

	if mhType == mh.SHA2_256 {
		prefix.Version = 0
	}

	return prefix
}

