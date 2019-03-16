
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460172030382080>

package corehttp

import (
	"net"
	"net/http"
	"runtime"
	"strconv"

	core "github.com/ipfs/go-ipfs/core"
)

//mutexfractionoption允许通过http设置runtime.setmutexprofilefraction
//使用参数“fraction”的post请求。
func MutexFractionOption(path string) ServeOption {
	return func(_ *core.IpfsNode, _ net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			asfr := r.Form.Get("fraction")
			if len(asfr) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			fr, err := strconv.Atoi(asfr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			log.Infof("Setting MutexProfileFraction to %d", fr)
			runtime.SetMutexProfileFraction(fr)
		})

		return mux, nil
	}
}

