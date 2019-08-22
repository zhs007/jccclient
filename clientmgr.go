package jccclient

// ClientMgr -
type ClientMgr struct {
	cfg     *Config
	Tags    map[string][]*Client
	NoTags  []*Client
	Clients map[string]*Client
}

// NewClientMgr - new clientmgr
func NewClientMgr(cfg *Config) *ClientMgr {
	mgr := &ClientMgr{
		cfg:     cfg,
		Tags:    make(map[string][]*Client),
		Clients: make(map[string]*Client),
	}

	mgr.init()

	return mgr
}

// init - init
func (mgr *ClientMgr) init() {
	for _, v := range mgr.cfg.Clients {
		cc := NewClient(v.ServAddr, v.Token)
		cc.cfg = mgr.cfg

		mgr.Clients[v.ServAddr] = cc

		if len(v.Tags) == 0 {
			mgr.NoTags = append(mgr.NoTags, cc)
		} else {
			for _, tv := range v.Tags {
				mgr.add2tag(tv, cc)
			}
		}
	}
}

// add2tag - add client to tag
func (mgr *ClientMgr) add2tag(tag string, c *Client) {
	_, isok := mgr.Tags[tag]
	if !isok {
		mgr.Tags[tag] = []*Client{}
	}

	mgr.Tags[tag] = append(mgr.Tags[tag], c)
}

// GetClient - get a client with hostname and tag
func (mgr *ClientMgr) GetClient(tag string, hostname string) *Client {
	if tag != "" {
		lst, isok := mgr.Tags[tag]
		if isok && len(lst) > 0 {
			return findClient(lst, hostname, mgr.cfg)
		}
	}

	return findClient(mgr.NoTags, hostname, mgr.cfg)
}
