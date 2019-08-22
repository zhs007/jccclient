package jccclient

import "time"

// HostInfo -
type HostInfo struct {
	Hostname       string
	TaskNums       int
	FailNums       int
	LastTime       int64
	LastFailTime   int64
	StartSleepTime int64
	SleepTime      int64
	MultiNums      int
}

// IsOK - is ok
func (hi *HostInfo) IsOK() bool {
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
	Hosts map[string]*HostInfo
}

// NewHostInfoCollection - new NewHostInfoCollection
func NewHostInfoCollection() *HostInfoCollection {
	return &HostInfoCollection{
		Hosts: make(map[string]*HostInfo),
	}
}

// Get - find a hostinfo
func (hic *HostInfoCollection) Get(hostname string) *HostInfo {
	hi, isok := hic.Hosts[hostname]
	if isok && hi != nil {
		return hi
	}

	return nil
}

// OnTaskStart - on task start
func (hic *HostInfoCollection) OnTaskStart(hostname string) {
	hi, isok := hic.Hosts[hostname]
	if !isok || hi == nil {
		hi = &HostInfo{
			Hostname: hostname,
		}

		hic.Hosts[hostname] = hi
	}

	hi.TaskNums++
	hi.LastTime = time.Now().Unix()
}

// OnTaskEnd - on task end
func (hic *HostInfoCollection) OnTaskEnd(hostname string, isfail bool, cfg *Config) {
	hi, isok := hic.Hosts[hostname]
	if isok && hi != nil {
		if isfail {
			ct := time.Now().Unix()

			hi.FailNums++

			if hi.LastFailTime > 0 && ct < hi.LastFailTime+int64(cfg.EffectiveFailTime) {
				hi.MultiNums++

				if hi.MultiNums >= cfg.MaxMultiFailTimes {
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
	}
}
