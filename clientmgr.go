package jccclient

// ClientMgr -
type ClientMgr struct {
	cfg     *Config
	db      *DB
	Tags    map[string][]*Client
	NoTags  []*Client
	Clients map[string]*Client
}

// NewClientMgr - new clientmgr
func NewClientMgr(cfg *Config) (*ClientMgr, error) {
	mgr := &ClientMgr{
		cfg:     cfg,
		Tags:    make(map[string][]*Client),
		Clients: make(map[string]*Client),
	}

	db, err := NewDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		return nil, err
	}

	mgr.db = db
	mgr.init()

	return mgr, nil
}

// init - init
func (mgr *ClientMgr) init() {
	for _, v := range mgr.cfg.Clients {
		cc := NewClient(mgr.db, v.ServAddr, v.Token)
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
