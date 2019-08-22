package jccclient

// findClient - find a client with hostname
func findClient(clients []*Client, hostname string, cfg *Config) *Client {
	var c *Client
	var chi *HostInfo

	for _, v := range clients {
		if c == nil {
			chi = c.Hosts.Get(hostname)
			if chi == nil {
				return v
			}

			if chi.IsOK() {
				c = v
			}
		} else {
			vhi := v.Hosts.Get(hostname)
			if vhi == nil {
				return v
			}

			if vhi.IsOK() {
				if vhi.LastTime < chi.LastTime {
					c = v
					chi = vhi
				}
			}
		}
	}

	return c
}
