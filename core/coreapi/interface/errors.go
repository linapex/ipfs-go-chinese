
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460165269164032>

package iface

import "errors"

var (
	ErrIsDir   = errors.New("this dag node is a directory")
	ErrNotFile = errors.New("this dag node is not a regular file")
	ErrOffline = errors.New("this action must be run in online mode, try running 'ipfs daemon' first")
)

