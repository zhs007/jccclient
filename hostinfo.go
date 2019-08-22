package jccclient

import (
	"context"
	"time"

	jccclientdbpb "github.com/zhs007/jccclient/dbpb"
)

// // HostInfo -
// type HostInfo struct {
// 	Hostname       string
// 	TaskNums       int
// 	FailNums       int
// 	LastTime       int64
// 	LastFailTime   int64
// 	StartSleepTime int64
// 	SleepTime      int64
// 	MultiNums      int
// }

// IsOK - is ok
func IsOK(hi *jccclientdbpb.HostInfo) bool {
	if hi.StartSleepTime > 0 {
		ct := time.Now().Unix()
		if ct > hi.StartSleepTime+hi.SleepTime {
			hi.StartSleepTime = 0

			return true
		}

		return false
	}

	return true
}

// HostInfoCollection -
type HostInfoCollection struct {
	Hosts    map[string]*jccclientdbpb.HostInfo
	servAddr string
	db       *DB
}

// NewHostInfoCollection - new NewHostInfoCollection
func NewHostInfoCollection(db *DB, servAddr string) *HostInfoCollection {
	return &HostInfoCollection{
		Hosts:    make(map[string]*jccclientdbpb.HostInfo),
		servAddr: servAddr,
		db:       db,
	}
}

// Get - find a hostinfo
func (hic *HostInfoCollection) Get(hostname string) *jccclientdbpb.HostInfo {
	hi, isok := hic.Hosts[hostname]
	if isok && hi != nil {
		return hi
	}

	return nil
}

// OnTaskStart - on task start
func (hic *HostInfoCollection) OnTaskStart(ctx context.Context, hostname string) {
	hi, isok := hic.Hosts[hostname]
	if !isok || hi == nil {
		hi = &jccclientdbpb.HostInfo{
			HostName: hostname,
		}

		hic.Hosts[hostname] = hi
	}

	hi.TaskNums++
	hi.LastTime = time.Now().Unix()

	if hic.db != nil {
		hic.db.UpdHostInfo(ctx, hic.servAddr, hostname, hi)
	}
}

// OnTaskEnd - on task end
func (hic *HostInfoCollection) OnTaskEnd(ctx context.Context, hostname string, isfail bool, cfg *Config) {
	hi, isok := hic.Hosts[hostname]
	if isok && hi != nil {
		if isfail {
			ct := time.Now().Unix()

			hi.FailNums++

			if hi.LastFailTime > 0 && ct < hi.LastFailTime+int64(cfg.EffectiveFailTime) {
				hi.MultiNums++

				if int(hi.MultiNums) >= cfg.MaxMultiFailTimes {
					hi.MultiNums = 0
					hi.StartSleepTime = ct

					if hi.SleepTime > 0 {
						hi.SleepTime *= 2
					} else {
						hi.SleepTime = int64(cfg.IgnoreTaskTime)
					}
				}
			}

			hi.LastFailTime = time.Now().Unix()
		} else {
			hi.MultiNums = 0
			hi.LastFailTime = 0
			hi.StartSleepTime = 0
			hi.SleepTime = 0
		}

		if hic.db != nil {
			hic.db.UpdHostInfo(ctx, hic.servAddr, hostname, hi)
		}
	}
}
