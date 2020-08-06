package jccclient

// JCCNode - jarviscrawlercore node
type JCCNode struct {
	Addr    string `json:"addr"`
	Version string `json:"version"`
	Err     string `json:"err"`
}

// JCCInfo - jarviscrawlercore infomation
type JCCInfo struct {
	Nodes []JCCNode `json:"nodes"`
}
