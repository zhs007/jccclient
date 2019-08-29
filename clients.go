package jccclient

import jccclientdbpb "github.com/zhs007/jccclient/dbpb"

// findClient - find a client with hostname
func findClient(tags *Tags, clients []*Client, hostname string, cfg *Config) *Client {
	var c *Client
	var chi *jccclientdbpb.HostInfo

	for _, v := range clients {
		if v.Running {
			continue
		}

		if tags != nil && !tags.IsMatch(v) {
			continue
		}

		if c == nil {
			chi = v.Hosts.Get(hostname)
			if chi == nil {
				return v
			}

			if IsOK(chi) {
				c = v
			}
		} else {
			vhi := v.Hosts.Get(hostname)
			if vhi == nil {
				return v
			}

			if IsOK(vhi) {
				if vhi.LastTime < chi.LastTime {
					c = v
					chi = vhi
				}
			}
		}
	}

	return c
}
