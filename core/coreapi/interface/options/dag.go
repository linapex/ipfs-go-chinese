
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460165722148864>

package options

import (
	"math"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
)

type DagPutSettings struct {
	InputEnc string
	Codec    uint64
	MhType   uint64
	MhLength int
}

type DagTreeSettings struct {
	Depth int
}

type DagPutOption func(*DagPutSettings) error
type DagTreeOption func(*DagTreeSettings) error

func DagPutOptions(opts ...DagPutOption) (*DagPutSettings, error) {
	options := &DagPutSettings{
		InputEnc: "json",
		Codec:    cid.DagCBOR,
		MhType:   math.MaxUint64,
		MhLength: -1,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func DagTreeOptions(opts ...DagTreeOption) (*DagTreeSettings, error) {
	options := &DagTreeSettings{
		Depth: -1,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type dagOpts struct{}

var Dag dagOpts

//inputenc是dag.put的一个选项，它指定
//数据。默认为“json”，大多数格式/编解码器支持“raw”
func (dagOpts) InputEnc(enc string) DagPutOption {
	return func(settings *DagPutSettings) error {
		settings.InputEnc = enc
		return nil
	}
}

//codec是dag.put的一个选项，它指定要使用的multicodec
//序列化对象。默认为cid.dagcbor（0x71）
func (dagOpts) Codec(codec uint64) DagPutOption {
	return func(settings *DagPutSettings) error {
		settings.Codec = codec
		return nil
	}
}

//hash是dag.put的一个选项，它指定要使用的多哈希设置
//散列对象时。默认值基于使用的编解码器
//（对于DAGCBOR，MH.SHA2U 256（0x12））。如果mhlen设置为-1，则默认长度为
//将使用哈希
func (dagOpts) Hash(mhType uint64, mhLen int) DagPutOption {
	return func(settings *DagPutSettings) error {
		settings.MhType = mhType
		settings.MhLength = mhLen
		return nil
	}
}

//深度是dag.tree的一个选项，它指定
//返回树。默认值为-1（无深度限制）
func (dagOpts) Depth(depth int) DagTreeOption {
	return func(settings *DagTreeSettings) error {
		settings.Depth = depth
		return nil
	}
}

