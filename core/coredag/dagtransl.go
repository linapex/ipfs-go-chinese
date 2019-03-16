
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460170960834560>

package coredag

import (
	"fmt"
	"io"

	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

//dagparser是用于将流解析为节点的函数
type DagParser func(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error)

//格式分析器用于将格式描述符映射到dagparsers
type FormatParsers map[string]DagParser

//inputencparsers用于将输入编码映射到格式化分析器。
type InputEncParsers map[string]FormatParsers

//defaultinputencparsers是在任何地方都使用的inputencparser
var DefaultInputEncParsers = InputEncParsers{
	"json":     defaultJSONParsers,
	"raw":      defaultRawParsers,
	"cbor":     defaultCborParsers,
	"protobuf": defaultProtobufParsers,
}

var defaultJSONParsers = FormatParsers{
	"cbor":     cborJSONParser,
	"dag-cbor": cborJSONParser,

	"protobuf": dagpbJSONParser,
	"dag-pb":   dagpbJSONParser,
}

var defaultRawParsers = FormatParsers{
	"cbor":     cborRawParser,
	"dag-cbor": cborRawParser,

	"protobuf": dagpbRawParser,
	"dag-pb":   dagpbRawParser,

	"raw": rawRawParser,
}

var defaultCborParsers = FormatParsers{
	"cbor":     cborRawParser,
	"dag-cbor": cborRawParser,
}

var defaultProtobufParsers = FormatParsers{
	"protobuf": dagpbRawParser,
	"dag-pb":   dagpbRawParser,
}

//ParseInputs使用DefaultInputencParsers分析IO.Reader
//输入IPLD节点实例的编码和格式
func ParseInputs(ienc, format string, r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	return DefaultInputEncParsers.ParseInputs(ienc, format, r, mhType, mhLen)
}

//addParser在给定输入编码和格式下添加dagparser
func (iep InputEncParsers) AddParser(ienc, format string, f DagParser) {
	m, ok := iep[ienc]
	if !ok {
		m = make(FormatParsers)
		iep[ienc] = m
	}

	m[format] = f
}

//parseinputs解析IO.reader，按输入编码和格式描述
//IPLD节点实例
func (iep InputEncParsers) ParseInputs(ienc, format string, r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	parsers, ok := iep[ienc]
	if !ok {
		return nil, fmt.Errorf("no input parser for %q", ienc)
	}

	parser, ok := parsers[format]
	if !ok {
		return nil, fmt.Errorf("no parser for format %q using input type %q", format, ienc)
	}

	return parser(r, mhType, mhLen)
}

