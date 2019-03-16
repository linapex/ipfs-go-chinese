
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:39</date>
//</624460163511750656>

package commands

import (
	"testing"
)

func TestCommandTree(t *testing.T) {
	printErrors := func(errs map[string][]error) {
		if errs == nil {
			return
		}
		t.Error("In Root command tree:")
		for cmd, err := range errs {
			t.Errorf("  In X command %s:", cmd)
			for _, e := range err {
				t.Errorf("    %s", e)
			}
		}
	}
	printErrors(Root.DebugValidate())
	printErrors(RootRO.DebugValidate())
}

