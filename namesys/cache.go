
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:43</date>
//</624460178376364032>

package namesys

import (
	"time"

	path "gx/ipfs/QmNYPETsdAu2uQ1k9q9S1jYEGURaLHV6cbYRSVFVRftpF8/go-path"
)

func (ns *mpns) cacheGet(name string) (path.Path, bool) {
	if ns.cache == nil {
		return "", false
	}

	ientry, ok := ns.cache.Get(name)
	if !ok {
		return "", false
	}

	entry, ok := ientry.(cacheEntry)
	if !ok {
//绝对不应该发生，纯粹为了理智
		log.Panicf("unexpected type %T in cache for %q.", ientry, name)
	}

	if time.Now().Before(entry.eol) {
		return entry.val, true
	}

	ns.cache.Remove(name)

	return "", false
}

func (ns *mpns) cacheSet(name string, val path.Path, ttl time.Duration) {
	if ns.cache == nil || ttl <= 0 {
		return
	}
	ns.cache.Add(name, cacheEntry{
		val: val,
		eol: time.Now().Add(ttl),
	})
}

type cacheEntry struct {
	val path.Path
	eol time.Time
}

