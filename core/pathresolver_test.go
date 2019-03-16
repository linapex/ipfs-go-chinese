
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:42</date>
//</624460173489999872>

package core_test

import (
	"testing"

	core "github.com/ipfs/go-ipfs/core"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	path "gx/ipfs/QmNYPETsdAu2uQ1k9q9S1jYEGURaLHV6cbYRSVFVRftpF8/go-path"
)

func TestResolveNoComponents(t *testing.T) {
	n, err := coremock.NewMockNode()
	if n == nil || err != nil {
		t.Fatal("Should have constructed a mock node", err)
	}

	_, err = core.Resolve(n.Context(), n.Namesys, n.Resolver, path.Path("/ipns/"))
	if err != path.ErrNoComponents {
		t.Fatal("Should error with no components (/ipns/).", err)
	}

	_, err = core.Resolve(n.Context(), n.Namesys, n.Resolver, path.Path("/ipfs/"))
	if err != path.ErrNoComponents {
		t.Fatal("Should error with no components (/ipfs/).", err)
	}

	_, err = core.Resolve(n.Context(), n.Namesys, n.Resolver, path.Path("/../.."))
	if err != path.ErrBadPath {
		t.Fatal("Should error with invalid path.", err)
	}
}

