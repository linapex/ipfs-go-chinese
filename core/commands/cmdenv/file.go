
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:38</date>
//</624460158965125120>

package cmdenv

import (
	"fmt"

	files "gx/ipfs/QmXWZCd8jfaHmt4UDSnjKmGcrQMw95bDGWqEeVLVJjoANX/go-ipfs-files"
)

//
func GetFileArg(it files.DirIterator) (files.File, error) {
	if !it.Next() {
		err := it.Err()
		if err == nil {
			err = fmt.Errorf("expected a file argument")
		}
		return nil, err
	}
	file := files.FileFromEntry(it)
	if file == nil {
		return nil, fmt.Errorf("file argument was nil")
	}
	return file, nil
}

