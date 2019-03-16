
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460171732586496>

package corehttp

import (
	"context"
	"net"
	"net/http"
	"strings"

	core "github.com/ipfs/go-ipfs/core"
	namesys "github.com/ipfs/go-ipfs/namesys"
	nsopts "github.com/ipfs/go-ipfs/namesys/opts"

	isd "gx/ipfs/QmZmmuAXgX73UQmX1jRKjTGmjzq24Jinqkq8vzkBtno4uX/go-is-domain"
)

//ipnshostnameoption重写传入请求，如果其host:header包含
//IPNS名称。
//重写的请求指向网关处理程序上解析的名称。
func IPNSHostnameOption() ServeOption {
	return func(n *core.IpfsNode, _ net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		childMux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithCancel(n.Context())
			defer cancel()

			host := strings.SplitN(r.Host, ":", 2)[0]
			if len(host) > 0 && isd.IsDomain(host) {
				name := "/ipns/" + host
				_, err := n.Namesys.Resolve(ctx, name, nsopts.Depth(1))
				if err == nil || err == namesys.ErrResolveRecursion {
					r.Header.Set("X-Ipns-Original-Path", r.URL.Path)
					r.URL.Path = name + r.URL.Path
				}
			}
			childMux.ServeHTTP(w, r)
		})
		return childMux, nil
	}
}

