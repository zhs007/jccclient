package jccclient

import "time"

// HostInfo -
type HostInfo struct {
	Hostname string
	TaskNums int
	ErrNums  int
	LastTime int64
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

// OnTaskFail - on task fail
func (hic *HostInfoCollection) OnTaskFail(hostname string) {
	hi, isok := hic.Hosts[hostname]
	if isok && hi != nil {
		hi.ErrNums++
	}
}
