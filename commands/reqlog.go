
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:37</date>
//</624460156532428800>

package commands

import (
	"sync"
	"time"
)

//reqlogentry是请求日志中的一个条目
type ReqLogEntry struct {
	StartTime time.Time
	EndTime   time.Time
	Active    bool
	Command   string
	Options   map[string]interface{}
	Args      []string
	ID        int

	log *ReqLog
}

//copy返回reqlogentry的副本
func (r *ReqLogEntry) Copy() *ReqLogEntry {
	out := *r
	out.log = nil
	return &out
}

//reqlog是请求的日志
type ReqLog struct {
	Requests []*ReqLogEntry
	nextID   int
	lock     sync.Mutex
	keep     time.Duration
}

//附录将一个条目添加到日志中
func (rl *ReqLog) AddEntry(rle *ReqLogEntry) {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	rl.nextID++
	rl.Requests = append(rl.Requests, rle)

	if rle == nil || !rle.Active {
		rl.maybeCleanup()
	}

	return
}

//ClearInactive删除过时的条目
func (rl *ReqLog) ClearInactive() {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	k := rl.keep
	rl.keep = 0
	rl.cleanup()
	rl.keep = k
}

func (rl *ReqLog) maybeCleanup() {
//只是偶尔做一次，否则可能
//成为性能问题
	if len(rl.Requests)%10 == 0 {
		rl.cleanup()
	}
}

func (rl *ReqLog) cleanup() {
	i := 0
	now := time.Now()
	for j := 0; j < len(rl.Requests); j++ {
		rj := rl.Requests[j]
		if rj.Active || rl.Requests[j].EndTime.Add(rl.keep).After(now) {
			rl.Requests[i] = rl.Requests[j]
			i++
		}
	}
	rl.Requests = rl.Requests[:i]
}

//setkeeptime设置一个持续时间，在此时间之后，一个条目将被视为不活动。
func (rl *ReqLog) SetKeepTime(t time.Duration) {
	rl.lock.Lock()
	defer rl.lock.Unlock()
	rl.keep = t
}

//报告生成请求日志中所有条目的副本
func (rl *ReqLog) Report() []*ReqLogEntry {
	rl.lock.Lock()
	defer rl.lock.Unlock()
	out := make([]*ReqLogEntry, len(rl.Requests))

	for i, e := range rl.Requests {
		out[i] = e.Copy()
	}

	return out
}

//finish将日志中的条目标记为finished
func (rl *ReqLog) Finish(rle *ReqLogEntry) {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	rle.Active = false
	rle.EndTime = time.Now()

	rl.maybeCleanup()
}

